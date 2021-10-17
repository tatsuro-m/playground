resource "google_service_account" "node_pool" {
  account_id   = "${local.app_prefix}-np"
  display_name = "gke node pool service account"
}

resource "google_project_iam_member" "role1" {
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:${google_service_account.node_pool.email}"
}

resource "google_project_iam_member" "role2" {
  role   = "roles/logging.logWriter"
  member = "serviceAccount:${google_service_account.node_pool.email}"
}

resource "google_service_account" "workload_identity" {
  account_id   = "${local.app_prefix}-main"
  display_name = "gke workload identity service account"
}


resource "google_container_cluster" "primary" {
  name     = "${local.app_prefix}-gke-cluster"
  location = var.default_region

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1

  network    = google_compute_network.main_vpc.id
  subnetwork = google_compute_subnetwork.main.id

  workload_identity_config {
    identity_namespace = "${data.google_project.project.project_id}.svc.id.goog"
  }
}

resource "google_container_node_pool" "main_node_pool" {
  name               = "${local.app_prefix}-main"
  location           = var.default_region
  cluster            = google_container_cluster.primary.name
  initial_node_count = 1

  autoscaling {
    min_node_count = 0
    max_node_count = 3
  }

  node_config {
    machine_type = "e2-micro"
    preemptible  = true

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
