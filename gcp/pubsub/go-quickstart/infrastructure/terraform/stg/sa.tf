resource "google_service_account" "main_topic" {
  account_id   = "${local.app_prefix}-main-topic"
  display_name = "pubsub main topic sa"
}
