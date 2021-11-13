resource "google_compute_network" "main_vpc" {
  name                    = "${local.app_prefix}-vpc-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "main" {
  name          = "${local.app_prefix}-gke-subnetwork"
  ip_cidr_range = "10.2.0.0/16"
  network       = google_compute_network.main_vpc.id
}

resource "google_compute_global_address" "default" {
  name = "ingress-test"
}
