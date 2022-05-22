resource "google_compute_instance_template" "tpl" {
  name         = "${local.app_prefix}-template"
  machine_type = "e2-medium"
  region       = "asia-northeast1"

  disk {
    source_image = "centos-cloud/centos-7"
    auto_delete  = true
    disk_size_gb = 35
    boot         = true
  }

  network_interface {
    network    = "default"
    subnetwork = "default"

    access_config {
      network_tier = "PREMIUM"
    }
  }

  metadata = {
    test = "true"
  }

  can_ip_forward = false
  metadata_startup_script = file("./script/start.sh")

  service_account {
    scopes = ["cloud-platform"]
    email = google_service_account.main.email
  }
}

resource "google_compute_instance_from_template" "test1" {
  name                     = "${local.app_prefix}-test1"
  source_instance_template = google_compute_instance_template.tpl.id
}

#resource "google_compute_instance_from_template" "test2" {
#  name                     = "${local.app_prefix}-test2"
#  source_instance_template = google_compute_instance_template.tpl.id
#}
