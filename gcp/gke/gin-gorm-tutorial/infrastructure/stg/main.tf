terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.77.0"
    }
  }
  backend "gcs" {
    bucket = "playground-tfstate-949181"
    prefix = "gin-gorm-tutorial/stg"
  }
}

provider "google" {
  project = "playground-318023"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}

resource "random_string" "random" {
  length  = 4
  number  = true
  upper   = false
  lower   = false
  special = false
}

locals {
  app_name     = "gin-gorm-tutorial"
  app_prefix = "stg-${local.app_name}"
}
