resource "random_string" "resource_name_suffix" {
  length  = 6
  number  = true
  upper   = false
  lower   = false
  special = false
}

resource "google_service_account" "a" {
  account_id   = "a-${random_string.resource_name_suffix.result}"
  display_name = "a"
  description  = "for test"
}
