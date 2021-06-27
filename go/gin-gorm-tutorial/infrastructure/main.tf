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
  project = "playground-318023"
  region  = "asia-northeast1"
}

resource "google_compute_network" "vpc_network" {
  name = "terraform-network"
}
