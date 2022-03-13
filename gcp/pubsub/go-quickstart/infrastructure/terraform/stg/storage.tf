resource "google_storage_bucket" "func_bucket1" {
  name     = "${local.app_prefix}-func-bucket1"
  location = "ASIA-NORTHEAST1"

  storage_class = "STANDARD"
  force_destroy = true
  uniform_bucket_level_access = true
}

data "archive_file" "function_archive" {
  type        = "zip"
  source_dir  = "./gcf/helloworld"
  output_path = "./output/zip/function.zip"
}

resource "google_storage_bucket_object" "archive1" {
  name   = "function.zip"
  bucket = google_storage_bucket.func_bucket1.name
  source = data.archive_file.function_archive.output_path
}
