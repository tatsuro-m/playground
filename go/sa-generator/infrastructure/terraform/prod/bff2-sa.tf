resource "google_service_account" "bff2" {
  account_id   = "${local.app_prefix}-bff2"
  display_name = "${local.app_prefix}-bff2"
  description  = "${local.app_prefix} の bff2 で利用する SA"
}
