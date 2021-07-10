resource "google_compute_instance" "default" {
  project      = google_project.my_project.project_id
  name         = "test"
  machine_type = "e2-micro"

  tags = [
    "foo",
  "bar"]

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    subnetwork = google_compute_subnetwork.basic.id
    access_config {
      // Ephemeral IP
    }
  }

}
