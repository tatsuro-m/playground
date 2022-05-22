terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.21.0"
    }
  }
  backend "gcs" {
    bucket = "playground-tfstate-949181"
    prefix = "gcsfuse/stg"
  }
}

provider "google" {
  project = "playground-318023"
  region = "asia-northeast1"
  zone   = "asia-northeast1-b"
}

variable "default_region" {
  type    = string
  default = "asia-northeast1"
}

variable "env" {
  type    = string
  default = "stg"
}

locals {
  app_name   = "gcsfuse"
  app_prefix = "stg-${local.app_name}"
}
