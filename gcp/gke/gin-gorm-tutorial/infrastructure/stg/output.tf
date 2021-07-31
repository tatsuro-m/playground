output "app_prefix" {
  value = local.app_prefix
}

output "repo-id" {
  value = google_artifact_registry_repository.main.id
}
