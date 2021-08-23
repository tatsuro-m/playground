resource "google_service_account" "crun" {
  account_id   = "${local.app_prefix}-crun"
  display_name = "cloud run default service account"
}

resource "google_project_iam_member" "role1" {
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:${google_service_account.crun.email}"
}

resource "google_cloud_run_service" "default" {
  name     = "${local.app_prefix}-default"
  location = var.default_region

  template {
    spec {
      service_account_name = google_service_account.crun.email
      containers {
        image = "asia-northeast1-docker.pkg.dev/playground-318023/stg-cloud-run-main/golang-api:v2"
      }
    }
  }

  metadata {
    annotations = {
      "autoscaling.knative.dev/minScale" = "0"
      "autoscaling.knative.dev/maxScale" = "10"
      "run.googleapis.com/client-name"   = "terraform"
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
