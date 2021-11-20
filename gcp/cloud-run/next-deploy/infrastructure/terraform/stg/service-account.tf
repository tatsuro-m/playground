resource "google_service_account" "frontend1" {
  account_id   = "${local.app_prefix}-frontend1"
  display_name = "next deploy frontend1 service account"
}

resource "google_service_account" "gin" {
  account_id   = "${local.app_prefix}-gin"
  display_name = "next deploy gin service account"
}
