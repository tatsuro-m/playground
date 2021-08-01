resource "google_service_account" "node_pool" {
  account_id   = "${local.app_prefix}-nodepool"
  display_name = "gke node pool service account"
}

resource "google_project_iam_member" "role1" {
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:${google_service_account.node_pool.email}"
}

resource "google_service_account" "workload_identity" {
  account_id   = "${local.app_prefix}-wi"
  display_name = "gke workload identity service account"
}

resource "google_project_iam_member" "role2" {
  role   = "roles/cloudsql.client"
  member = "serviceAccount:${google_service_account.workload_identity.email}"
}

// サービスアカウントをリソースとして使うので、 ksa からアクセスできるように bind する
resource "google_service_account_iam_member" "admin-account-iam" {
  service_account_id = google_service_account.workload_identity.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "serviceAccount:${data.google_project.project.project_id}.svc.id.goog[${local.app_name}/main-ksa]"
}

resource "google_container_cluster" "primary" {
  name     = "${local.app_prefix}-gke-cluster"
  location = "asia-northeast1"

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
  location   = "asia-northeast1"
  cluster    = google_container_cluster.primary.name
  node_count = 1

  node_config {
    machine_type = "e2-medium"

    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.node_pool.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
