resource "google_service_account" "rpcserver" {
  account_id   = "${local.app_prefix}-rpcserver"
  display_name = "${local.app_prefix}-rpcserver"
  description  = "${local.app_prefix} の rpcserver で利用する SA"
}
