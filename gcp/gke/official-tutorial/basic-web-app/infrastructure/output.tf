output "app_name" {
  value = local.app_name
}

output "repo-id" {
  value = google_artifact_registry_repository.my-repo.id
}
