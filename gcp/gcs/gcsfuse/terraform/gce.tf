resource "google_compute_instance_template" "tpl" {
  name         = "${local.app_prefix}-template"
  machine_type = "e2-medium"
  region       = "asia-northeast1"
  tags         = [
    "es"
  ]

  disk {
    source_image = "centos-cloud/centos-7"
    auto_delete  = true
    disk_size_gb = 30
    boot         = true
  }

  disk {
    source = google_compute_disk.data.name
    auto_delete = true
    boot = false
  }

  network_interface {
    network    = google_compute_network.vpc_network.id
    subnetwork = google_compute_subnetwork.private_1.id

    access_config {
      network_tier = "PREMIUM"
    }
  }

  metadata = {
    init-script = file("./script/init-script.sh")
    es-config = file("./script/elasticsearch.yml")
  }

  metadata_startup_script = file("./script/startup-script.sh")

  service_account {
    scopes = ["cloud-platform"]
    email  = google_service_account.main.email
  }
}

resource "google_compute_disk" "data" {
  name  = "${local.app_prefix}-data"
  size  = 10
  type  = "pd-ssd"
  zone  = "asia-northeast1-b"
}


resource "google_compute_instance_from_template" "test1" {
  name                     = "${local.app_prefix}-test1"
  source_instance_template = google_compute_instance_template.tpl.id

  network_interface {
    network    = google_compute_network.vpc_network.id
    subnetwork = google_compute_subnetwork.private_1.id
    network_ip = "10.0.0.12"

    access_config {
      network_tier = "PREMIUM"
    }
  }
}

#resource "google_compute_instance_from_template" "test2" {
#  name                     = "${local.app_prefix}-test2"
#  source_instance_template = google_compute_instance_template.tpl.id
#}
