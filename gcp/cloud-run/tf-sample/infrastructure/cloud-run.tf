resource "google_cloud_run_service" "default" {
  name     = "${local.app_prefix}-default"
  location = "asia-northeast1"

  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
