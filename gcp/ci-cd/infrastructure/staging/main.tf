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
  region = "asia-northeast1"
  zone   = "asia-northeast1-b"
}

variable "billing_account_id" {}

resource "random_string" "random" {
  length  = 4
  number  = true
  upper   = false
  lower   = false
  special = false
}

locals {
  app_name = "ci-cd-lesson"
  project_name = "stg-${local.app_name}"
}

resource "google_project" "my_project" {
  name                = local.project_name
  project_id          = "${local.project_name}-${random_string.random.result}"
  billing_account     = var.billing_account_id
  auto_create_network = false
}
