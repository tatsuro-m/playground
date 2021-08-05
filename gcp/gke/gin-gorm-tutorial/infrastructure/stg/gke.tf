resource "google_service_account" "node_pool" {
  account_id   = "${local.app_prefix}-nodepool"
  display_name = "gke node pool service account"
}

resource "google_project_iam_member" "role1" {
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:${google_service_account.node_pool.email}"
}

resource "google_service_account" "workload_identity" {
  account_id   = "${local.app_prefix}-main"
  display_name = "gke workload identity service account"
}

resource "google_project_iam_member" "role2" {
  role   = "roles/cloudsql.client"
  member = "serviceAccount:${google_service_account.workload_identity.email}"
}

// サービスアカウントをリソースとして使うので、 ksa からアクセスできるように bind する
resource "google_service_account_iam_member" "bind1" {
  service_account_id = google_service_account.workload_identity.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "serviceAccount:${data.google_project.project.project_id}.svc.id.goog[${local.app_name}/main-ksa]"
}

resource "google_service_account" "external_secret" {
  account_id   = "${local.app_prefix}-ex-sec"
  display_name = "gke external secret service account"
}

resource "google_service_account_iam_member" "bind2" {
  service_account_id = google_service_account.external_secret.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "serviceAccount:${data.google_project.project.project_id}.svc.id.goog[${local.app_name}/kubernetes-external-secrets]"
}

resource "google_secret_manager_secret_iam_member" "bind3" {
  secret_id = google_secret_manager_secret.dsn.secret_id
  role      = "roles/secretmanager.secretAccessor"
  member    = "serviceAccount:${google_service_account.external_secret.email}"
}


resource "google_container_cluster" "primary" {
  name     = "${local.app_prefix}-gke-cluster"
  location = var.default_region

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1

  workload_identity_config {
    identity_namespace = "${data.google_project.project.project_id}.svc.id.goog"
  }
}

resource "google_container_node_pool" "main_node_pool" {
  name       = "${local.app_prefix}-main"
  location   = var.default_region
  cluster    = google_container_cluster.primary.name
  node_count = 1

  node_config {
    machine_type = "e2-medium"

    workload_metadata_config {
      node_metadata = "GKE_METADATA_SERVER"
    }

    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.node_pool.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
