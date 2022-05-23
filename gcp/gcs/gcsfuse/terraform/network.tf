resource "google_compute_network" "vpc_network" {
  name = "${local.app_prefix}-vpc-1"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "private_1" {
  name          = "${local.app_prefix}private-2"
  ip_cidr_range = "10.0.0.0/20"
  region        = "asia-northeast1"
  network       = google_compute_network.vpc_network.id
}

resource "google_compute_firewall" "allow" {
  name    = "${local.app_prefix}-allow"
  network = google_compute_network.vpc_network.id

  allow {
    protocol = "tcp"
    ports    = ["80", "443", "22"]
  }

  target_tags = [
    "es"
  ]
}
