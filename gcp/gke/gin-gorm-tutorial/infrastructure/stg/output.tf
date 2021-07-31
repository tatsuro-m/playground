output "app_prefix" {
  value = local.app_prefix
}

output "repo-url" {
  value = "${google_artifact_registry_repository.main.location}-docker.pkg.dev/${google_artifact_registry_repository.main.project}/${google_artifact_registry_repository.main.repository_id}"
}

output "cluster_name" {
  value = google_container_cluster.primary.name
}
