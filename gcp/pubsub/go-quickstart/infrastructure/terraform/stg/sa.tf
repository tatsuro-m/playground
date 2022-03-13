resource "google_service_account" "main_topic" {
  account_id   = "${local.app_prefix}-main-topic"
  display_name = "pubsub main topic sa"
}

resource "google_service_account_key" "main_topic_key" {
  service_account_id = google_service_account.main_topic.name
}

resource "local_file" main_topic_key {
  filename             = "./output/secrets/${google_service_account.main_topic.name}.json"
  content              = base64decode(google_service_account_key.main_topic_key.private_key)
  file_permission      = "0600"
  directory_permission = "0755"
}

resource "google_project_iam_member" "mt_pubsub_admin" {
  project = "playground-318023"
  role    = "roles/pubsub.admin"
  member  = "serviceAccount:${google_service_account.main_topic.email}"
}

resource "google_service_account" "gcf1" {
  account_id   = "${local.app_prefix}-gcf-runtime"
}
