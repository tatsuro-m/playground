resource "google_storage_bucket" "a" {
  name          = "${local.app_name}-test-bucket-1"
  location      = "ASIA-NORTHEAST1"
  force_destroy = true

  uniform_bucket_level_access = true
}

//resource "google_storage_bucket" "b" {
//  project = google_project.my_project.project_id
//  name          = "test-bucket-${random_string.resource_name_suffix.result}-2"
//  location      = "ASIA-NORTHEAST1"
//  force_destroy = true
//
//  uniform_bucket_level_access = true
//}
