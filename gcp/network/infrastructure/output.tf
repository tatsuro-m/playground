output "project_id" {
  value = google_project.my_project.project_id
}

output "gce_instance_public_ip" {
  value = google_compute_instance.default.network_interface.0.access_config.0.nat_ip
}

output "db_public_ip" {
  value = google_sql_database_instance.master.public_ip_address
}

output "db_private_ip" {
  value = google_sql_database_instance.master.private_ip_address
}
