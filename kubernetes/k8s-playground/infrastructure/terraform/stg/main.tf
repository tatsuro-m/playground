terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.6.0"
    }
  }
  backend "gcs" {
    bucket = "playground-tfstate-949181"
    prefix = "k8s-pg/stg"
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
  app_name   = "k8s-pg"
  app_prefix = "stg-${local.app_name}"
}
