resource "google_service_account" "main" {
  account_id   = "${local.app_prefix}-main"
  display_name = "バケットにアクセスする際のテスト用"
}

resource "google_service_account_key" "main" {
  service_account_id = google_service_account.main.name
}

resource "local_file" "my-account-key" {
  filename             = "./output/secrets/gcp-sa-key.json"
  content              = base64decode(google_service_account_key.main.private_key)
  file_permission      = "0600"
  directory_permission = "0755"
}

resource "google_storage_bucket_iam_member" "bind1" {
  bucket = google_storage_bucket.test1.name
  role = "roles/storage.admin"
  member = "serviceAccount:${google_service_account.main.email}"
}
