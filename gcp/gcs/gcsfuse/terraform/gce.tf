resource "google_compute_instance_template" "tpl" {
  name         = "${local.app_prefix}-template"
  machine_type = "e2-small"
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

  metadata_startup_script = <<-EOT
#!/bin/bash

sudo tee /etc/yum.repos.d/gcsfuse.repo > /dev/null <<EOF
[gcsfuse]
name=gcsfuse (packages.cloud.google.com)
baseurl=https://packages.cloud.google.com/yum/repos/gcsfuse-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=0
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
       https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF

sudo yum install gcsfuse -y

mkdir -p /root/gcsfuse/test1
gcsfuse stg-gcsfuse-test1 /root/gcsfuse/test1

mkdir -p /root/gcsfuse/test2
gcsfuse stg-gcsfuse-test2 /root/gcsfuse/test2

EOT

  service_account {
    scopes = ["cloud-platform"]
    email = google_service_account.main.email
  }
}

resource "google_compute_instance_from_template" "test1" {
  name                     = "${local.app_prefix}-test1"
  source_instance_template = google_compute_instance_template.tpl.id
}
