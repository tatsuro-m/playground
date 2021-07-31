resource "google_sql_database_instance" "master" {
  name                = "${local.app_prefix}-master-instance"
  database_version    = "POSTGRES_13"
  deletion_protection = false

  settings {
    tier = "db-f1-micro"

    ip_configuration {
      ipv4_enabled = true
    }
  }
}

variable "POSTGRES_DB_NAME" {}
resource "google_sql_database" "database" {
  name     = "${local.app_prefix}-${var.POSTGRES_DB_NAME}"
  instance = google_sql_database_instance.master.name
}

variable "POSTGRES_PASSWORD" {}
resource "google_sql_user" "users" {
  name     = "postgres"
  instance = google_sql_database_instance.master.name
  password = var.POSTGRES_PASSWORD
}
