resource "google_cloud_run_service" "default" {
  name     = "${local.app_prefix}-default"
  location = var.default_region

  template {
    spec {
      containers {
        image = "asia-northeast1-docker.pkg.dev/playground-318023/stg-cloud-run-main/golang-api:v2"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location = google_cloud_run_service.default.location
  project  = google_cloud_run_service.default.project
  service  = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}


resource "google_artifact_registry_repository" "main" {
  provider = google-beta
  project  = "playground-318023"

  location      = var.default_region
  repository_id = "${local.app_prefix}-main"
  description   = "main golang repository"
  format        = "DOCKER"
}
