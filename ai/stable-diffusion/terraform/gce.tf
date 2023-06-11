resource "google_compute_instance" "main" {
  name         = "${local.app_prefix}-main"
  machine_type = "n1-standard-4"
  zone         = "asia-northeast1-a"

  boot_disk {
    auto_delete = true
    device_name = "instance-1"

    initialize_params {
      image = "projects/ml-images/global/images/c0-deeplearning-common-cu113-v20230501-debian-10"
      size  = 50
      type  = "pd-balanced"
    }

    mode = "READ_WRITE"
  }

  scheduling {
    preemptible       = true
    automatic_restart = false
    provisioning_model = "SPOT"
    on_host_maintenance = "TERMINATE"
  }

  guest_accelerator {
    count = 1
    type  = "nvidia-tesla-t4"
  }

  network_interface {
    network = "default"
    access_config {
    }
  }
}
