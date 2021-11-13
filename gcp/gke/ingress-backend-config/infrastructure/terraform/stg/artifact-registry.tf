resource "google_artifact_registry_repository" "frontA" {
  provider = google-beta
  project  = "playground-318023"

  location      = var.default_region
  repository_id = "${local.app_prefix}-frontB"
  description   = "web front end A"
  format        = "DOCKER"
}

resource "google_artifact_registry_repository" "frontB" {
  provider = google-beta
  project  = "playground-318023"

  location      = var.default_region
  repository_id = "${local.app_prefix}-frontB"
  description   = "web front end B"
  format        = "DOCKER"
}
