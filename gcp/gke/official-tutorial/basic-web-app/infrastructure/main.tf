terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.73.0"
    }
  }
  backend "gcs" {
    bucket = "playground-tfstate-949181"
    prefix = "gke-basic-tutorial/stg"
  }
}

provider "google" {
  project = "playground-318023"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-b"
}

resource "random_string" "random" {
  length  = 4
  number  = true
  upper   = false
  lower   = false
  special = false
}

locals {
  app_name     = "gke-basic-tutorial"
  project_name = "stg-${local.app_name}"
}
