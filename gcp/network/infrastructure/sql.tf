resource "google_sql_database_instance" "master" {
  project          = google_project.my_project.project_id
  name             = "master-instance"
  database_version = "POSTGRES_13"
  deletion_protection = false

  settings {
    tier = "db-f1-micro"

    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.vpc_network.id
    }
  }
}

variable "POSTGRES_DB_NAME" {}
resource "google_sql_database" "database" {
  project  = google_project.my_project.project_id
  name     = var.POSTGRES_DB_NAME
  instance = google_sql_database_instance.master.name
}

variable "POSTGRES_PASSWORD" {}
resource "google_sql_user" "users" {
  project  = google_project.my_project.project_id
  name     = "postgres"
  instance = google_sql_database_instance.master.name
  password = var.POSTGRES_PASSWORD
}
