resource "google_artifact_registry_repository" "nginx" {
  provider = google-beta
  project  = "playground-318023"

  location      = var.default_region
  repository_id = "${local.app_prefix}-nginx-proxy"
  description   = "nginx reverse proxy repository"
  format        = "DOCKER"
}

resource "google_artifact_registry_repository" "frontend" {
  provider = google-beta
  project  = "playground-318023"

  location      = var.default_region
  repository_id = "${local.app_prefix}-frontend"
  description   = "web front end"
  format        = "DOCKER"
}
