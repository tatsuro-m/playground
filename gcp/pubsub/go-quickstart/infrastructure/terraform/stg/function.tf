resource "google_cloudfunctions_function" "function" {
  name        = "${local.app_prefix}-test1"
  description = "My function"
  runtime     = "go116"

  available_memory_mb   = 128
  source_archive_bucket = google_storage_bucket.func_bucket1.name
  source_archive_object = google_storage_bucket_object.archive1.name
  service_account_email = google_service_account.gcf1.email

  event_trigger {
    event_type = "google.pubsub.topic.publish"
    resource   = google_pubsub_topic.my_topic.name
  }
  entry_point = "HelloPubSub"
}
