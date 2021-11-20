resource "google_service_account" "next_deploy_frontend1" {
  account_id   = "${local.app_prefix}-frontend1"
  display_name = "next deploy frontend1 service account"
}

resource "google_project_iam_member" "role1" {
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:${google_service_account.next_deploy_frontend1.email}"
}
