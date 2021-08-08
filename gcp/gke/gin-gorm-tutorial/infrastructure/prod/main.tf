terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.77.0"
    }
  }
  backend "gcs" {
    bucket = "playground-tfstate-949181"
    prefix = "gin-gorm-tutorial/prod"
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

data "google_project" "project" {}

variable "default_region" {
  type    = string
  default = "asia-northeast1"
}

locals {
  app_name   = "gin-gorm-tutorial"
  app_prefix = "prod-${local.app_name}"
}
