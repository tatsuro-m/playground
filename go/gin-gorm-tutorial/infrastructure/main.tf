terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.5.0"
    }
  }
  backend "gcs" {
    bucket  = "playground-tfstate-949181"
    prefix  = "gin-gorm-tutorial/dev"
  }
}

provider "google" {
  region  = "asia-northeast1"
}

variable "billing_account_id" {}

resource "google_project" "my_project" {
  name       = "gin-gorm-tutorial"
  project_id = "gin-gorm-tutorial-2348792381"
  billing_account = var.billing_account_id
}

resource "google_compute_network" "vpc_network" {
  project = google_project.my_project.project_id
  name = "terraform-network"
}
