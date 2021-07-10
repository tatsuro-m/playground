resource "google_compute_network" "vpc_network" {
  project                 = google_project.my_project.project_id
  name                    = "vpc-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "basic" {
  project       = google_project.my_project.project_id
  name          = "test-subnetwork"
  ip_cidr_range = "10.2.0.0/16"
  network       = google_compute_network.vpc_network.id
}

resource "google_compute_firewall" "gce" {
  project = google_project.my_project.project_id
  name    = "gce-firewall"
  network = google_compute_network.vpc_network.name

  allow {
    protocol = "tcp"
    ports = ["22"]
  }
}
