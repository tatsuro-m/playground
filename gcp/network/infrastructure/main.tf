terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "3.73.0"
    }
  }
  backend "gcs" {
    bucket = "playground-tfstate-949181"
    prefix = "network/dev"
  }
}

provider "google" {
  region = "asia-northeast1"
}

variable "billing_account_id" {}

resource "google_project" "my_project" {
  name            = "network-lesson"
  project_id      = "network-lesson-234879292381"
  billing_account = var.billing_account_id
}
