terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.5.0"
    }
  }
  backend "gcs" {
    bucket = "playground-tfstate-949181"
    prefix = "gin-gorm-tutorial/dev"
  }
}

provider "google" {
  region = "asia-northeast1"
}

variable "billing_account_id" {}

resource "google_project" "my_project" {
  name            = "gin-gorm-tutorial"
  project_id      = "gin-gorm-tutorial-2348792381"
  billing_account = var.billing_account_id
}

resource "google_sql_database_instance" "master" {
  project          = google_project.my_project.project_id
  name             = "master-instance"
  database_version = "POSTGRES_13"

  settings {
    tier = "db-f1-micro"

    ip_configuration {
      authorized_networks {
        name  = "all network access"
        value = "0.0.0.0/0"
      }
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

