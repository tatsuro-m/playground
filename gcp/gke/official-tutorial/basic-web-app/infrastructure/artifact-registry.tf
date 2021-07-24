resource "google_artifact_registry_repository" "my-repo" {
  provider = google-beta
  project  = "playground-318023"

  location      = "asia-northeast1"
  repository_id = "${local.project_name}-hello-repo"
  description   = "example docker repository"
  format        = "DOCKER"
}
