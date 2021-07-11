resource "google_storage_bucket" "a" {
  project = google_project.my_project.project_id
  name          = "test-bucket-${random_string.resource_name_suffix.result}"
  location      = "ASIA-NORTHEAST1"
  force_destroy = true

  uniform_bucket_level_access = true
}
