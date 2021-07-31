resource "google_service_account" "gke" {
  account_id   = "${local.app_prefix}-main"
  display_name = "gke service account"
}

resource "google_project_iam_member" "sample1" {
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:${google_service_account.gke.email}"
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
    machine_type = "e2-micro"

    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.gke.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
