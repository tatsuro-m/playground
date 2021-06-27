output "project_id" {
  value = google_project.my_project.project_id
}

output "db_public_ip" {
  value = google_sql_database_instance.master.public_ip_address
}
