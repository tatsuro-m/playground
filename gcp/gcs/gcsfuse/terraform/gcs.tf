resource "google_storage_bucket" "test1" {
  name          = "${local.app_prefix}-test1"
  location      = "ASIA-NORTHEAST1"
  force_destroy = true

  uniform_bucket_level_access = true
}
