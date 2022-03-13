resource "google_pubsub_topic" "my_topic" {
  name = "${local.app_prefix}-my-topic"
  schema_settings {
    schema = google_pubsub_schema.main.id
    encoding = "JSON"
  }

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

  enable_message_ordering = false
}

resource "google_pubsub_schema" "main" {
  name = "${local.app_prefix}-main"
  type = "AVRO"

  definition = <<EOF
{
  "type" : "record",
  "name" : "Avro",
  "fields" : [
    {
      "name" : "StringField",
      "type" : "string"
    },
    {
      "name" : "IntField",
      "type" : "int"
    }
  ]
}
EOF
}
