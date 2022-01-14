resource "google_service_account" "default_node_pool" {
  account_id   = "${local.app_prefix}-default-np"
  display_name = "gke default node pool service account"
}

resource "google_project_iam_member" "bind1" {
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:${google_service_account.default_node_pool.email}"
}

resource "google_project_iam_member" "bind2" {
  role   = "roles/logging.logWriter"
  member = "serviceAccount:${google_service_account.default_node_pool.email}"
}

resource "google_service_account" "workload_identity" {
  account_id   = "${local.app_prefix}-wi"
  display_name = "gke workload identity service account"
}

// サービスアカウントをリソースとして使うので、 ksa からアクセスできるように bind する
resource "google_service_account_iam_member" "bind3" {
  service_account_id = google_service_account.workload_identity.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "serviceAccount:${data.google_project.project.project_id}.svc.id.goog[${local.app_name}/main-ksa]"
}

resource "google_service_account" "external_secret" {
  account_id   = "${local.app_prefix}-ex-sec"
  display_name = "gke external secret service account"
}

resource "google_service_account_iam_member" "bind4" {
  service_account_id = google_service_account.external_secret.name
  role               = "roles/iam.workloadIdentityUser"
  member             = "serviceAccount:${data.google_project.project.project_id}.svc.id.goog[${local.app_name}/external-secrets-ksa]"
}
