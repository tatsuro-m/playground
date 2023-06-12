resource "google_compute_instance" "main" {
  name         = "${local.app_prefix}-main"
  machine_type = "n1-standard-4"
  zone         = "asia-northeast1-a"

  boot_disk {
    auto_delete = true
    device_name = "instance-1"

    initialize_params {
      image = "projects/ubuntu-os-cloud/global/images/ubuntu-minimal-2204-jammy-v20230605"
      size  = 50
      type  = "pd-balanced"
    }

    mode = "READ_WRITE"
  }

  can_ip_forward      = false
  deletion_protection = false
  enable_display      = false

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
