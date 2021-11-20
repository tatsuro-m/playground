resource "google_service_account" "next_deploy_frontend1" {
  account_id   = "${local.app_prefix}-frontend1"
  display_name = "next deploy frontend1 service account"
}
