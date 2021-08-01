output "app_prefix" {
  value = local.app_prefix
}

output "image_repo_url" {
  value = "${google_artifact_registry_repository.main.location}-docker.pkg.dev/${google_artifact_registry_repository.main.project}/${google_artifact_registry_repository.main.repository_id}"
}

output "cluster_name" {
  value = google_container_cluster.primary.name
}

output "db_public_ip" {
  value = google_sql_database_instance.master.ip_address.0.ip_address
}

output "db_connection_name" {
  value = google_sql_database_instance.master.connection_name
}

output "db_name" {
  value = google_sql_database.database.name
}

output "dsn_secret_name" {
  value = google_secret_manager_secret.dsn.secret_id
}
