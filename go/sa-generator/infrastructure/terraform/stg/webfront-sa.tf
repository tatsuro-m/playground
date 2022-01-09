resource "google_service_account" "webfront" {
  account_id   = "${local.app_prefix}-webfront"
  display_name = "${local.app_prefix}-webfront"
  description  = "${local.app_prefix} の webfront で利用する SA"
}
