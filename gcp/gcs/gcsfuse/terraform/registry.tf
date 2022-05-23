resource "google_artifact_registry_repository" "es" {
  provider = google-beta
  project = "playground-318023"

  location = "asia-northeast1"
  repository_id = "${local.app_prefix}-es"
  description = "example docker repository"
  format = "DOCKER"
}

resource "google_artifact_registry_repository_iam_member" "member1" {
  provider = google-beta
  project = google_artifact_registry_repository.es.project

  location = google_artifact_registry_repository.es.location
  repository = google_artifact_registry_repository.es.name
  role = "roles/artifactregistry.repoAdmin"
  member = "serviceAccount:${google_service_account.main.email}"
}
