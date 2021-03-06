resource "google_artifact_registry_repository" "main" {
  provider = google-beta
  project  = "playground-318023"

  location      = var.default_region
  repository_id = "${local.app_prefix}-main"
  description   = "main golang repository"
  format        = "DOCKER"
}
