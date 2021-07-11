resource "google_storage_bucket" "a" {
  name          = "${local.app_name}-test-bucket-a"
  location      = "ASIA-NORTHEAST1"
  force_destroy = true

  uniform_bucket_level_access = true
}

resource "google_storage_bucket" "b" {
  name          = "${local.app_name}-test-bucket-b"
  location      = "ASIA-NORTHEAST1"
  force_destroy = true

  uniform_bucket_level_access = true
}
