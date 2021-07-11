terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.73.0"
    }
  }
  backend "gcs" {
    bucket = "playground-tfstate-949181"
    prefix = "ci-cd/stg"
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
  app_name     = "ci-cd-lesson"
  project_name = "stg-${local.app_name}"
}
