resource "google_pubsub_topic" "my_topic" {
  name = "${local.app_prefix}-my-topic"

  labels = {
    foo = "bar"
  }

  message_retention_duration = "86600s"
}

# pull 型のサブスクリプション
resource "google_pubsub_subscription" "my_topic_sub1" {
  name  = "${local.app_prefix}-my-topic-sub1"
  topic = google_pubsub_topic.my_topic.name

  labels = {
    foo = "bar"
  }

  # 20 minutes
  message_retention_duration = "1200s"
  retain_acked_messages      = true

  ack_deadline_seconds = 20

  expiration_policy {
    ttl = "300000.5s"
  }
  retry_policy {
    minimum_backoff = "10s"
  }

  enable_message_ordering    = false
}
