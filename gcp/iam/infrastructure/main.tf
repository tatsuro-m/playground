terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.73.0"
    }
  }
  backend "gcs" {
    bucket = "playground-tfstate-949181"
    prefix = "iam/dev"
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


resource "google_project" "my_project" {
  name                = "iam-lesson"
  project_id          = "iam-lesson-${random_string.random.result}"
  billing_account     = var.billing_account_id
  auto_create_network = false
}
