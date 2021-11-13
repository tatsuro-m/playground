resource "google_artifact_registry_repository" "front1" {
  provider = google-beta
  project  = "playground-318023"

  location      = var.default_region
  repository_id = "${local.app_prefix}-front1"
  description   = "web front end 1"
  format        = "DOCKER"
}

resource "google_artifact_registry_repository" "front2" {
  provider = google-beta
  project  = "playground-318023"

  location      = var.default_region
  repository_id = "${local.app_prefix}-front2"
  description   = "web front end 2"
  format        = "DOCKER"
}
