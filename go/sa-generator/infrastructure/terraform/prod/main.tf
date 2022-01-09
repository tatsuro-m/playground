terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.5.0"
    }
  }
  backend "gcs" {
    bucket = "playground-tfstate-949181"
    prefix = "sa-generator/prod"
  }
}

provider "google" {
  project = "playground-318023"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}

data "google_project" "project" {}

variable "default_region" {
  type    = string
  default = "asia-northeast1"
}

locals {
  app_name   = "sa-generator"
  app_prefix = "prod-${local.app_name}"
}
