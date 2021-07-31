output "app_prefix" {
  value = local.app_prefix
}

output "repo-id" {
  value = google_artifact_registry_repository.main.id
}

output "cluster_name" {
  value = google_container_cluster.primary.name
}
