resource "google_service_account" "ssh-account" {
  project = google_project.my_project.project_id
  account_id   = "gce-ssh-account"
  display_name = "ssh-account"
}

resource "google_compute_instance" "default" {
  project      = google_project.my_project.project_id
  name         = "test"
  machine_type = "e2-micro"
  allow_stopping_for_update = true

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

  service_account {
    email  = google_service_account.ssh-account.email
    scopes = ["cloud-platform"]
  }
}
