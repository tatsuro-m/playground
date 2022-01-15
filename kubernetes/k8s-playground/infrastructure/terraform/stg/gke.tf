resource "google_container_cluster" "main" {
  #  環境ごとにクラスターを分けるのではなく、 namespace で分離する方法を取る
  #  実プロダクトなら、おそらくクラスターレベルで分けるのが良い
  name     = "${local.app_name}-main"
  location = "asia-northeast1-a"
  node_locations = [
    "asia-northeast1-b"
  ]

  # We can't create a cluster with no node pool defined, but we want to only use
  # separately managed node pools. So we create the smallest possible default
  # node pool and immediately delete it.
  remove_default_node_pool = true
  initial_node_count       = 1

  network    = google_compute_network.main_vpc.id
  subnetwork = google_compute_subnetwork.main.id

  workload_identity_config {
    workload_pool = "${data.google_project.project.project_id}.svc.id.goog"
  }
}

resource "google_container_node_pool" "main_default_node_pool" {
  name       = "${local.app_prefix}-default"
  cluster    = google_container_cluster.main.name
  node_count = 1

  node_config {
    machine_type = "e2-micro"
    preemptible  = false
    disk_size_gb = 10
    disk_type    = "pd-standard"

    workload_metadata_config {
      mode = "GKE_METADATA"
    }

    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
    service_account = google_service_account.default_node_pool.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
