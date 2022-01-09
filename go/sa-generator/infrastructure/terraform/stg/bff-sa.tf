resource "google_service_account" "bff" {
  account_id   = "${local.app_prefix}-bff"
  display_name = "${local.app_prefix}-bff"
  description  = "${local.app_prefix} の bff で利用する SA"
}
