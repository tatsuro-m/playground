resource "google_artifact_registry_repository" "frontend1" {
  provider = google-beta
  project  = "playground-318023"

  location      = var.default_region
  repository_id = "${local.app_prefix}-frontend1"
  description   = "web frontend 1"
  format        = "DOCKER"
}
