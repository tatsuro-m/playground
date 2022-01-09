resource "google_service_account" "adminfront" {
  account_id   = "${local.app_prefix}-adminfront"
  display_name = "${local.app_prefix}-adminfront"
  description  = "${local.app_prefix} の adminfront で利用する SA"
}
