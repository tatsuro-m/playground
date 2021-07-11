resource "random_string" "iam_suffix" {
  length  = 6
  number  = true
  upper   = false
  lower   = false
  special = false
}

resource "google_service_account" "a" {
  project = google_project.my_project.project_id
  account_id   = "a-${random_string.iam_suffix.result}"
  display_name = "a"
  description = "for test"
}

resource "google_service_account" "b" {
  project = google_project.my_project.project_id
  account_id   = "b-${random_string.iam_suffix.result}"
  display_name = "b"
  description = "for test"
}

resource "google_service_account" "c" {
  project = google_project.my_project.project_id
  account_id   = "c-${random_string.iam_suffix.result}"
  display_name = "c"
  description = "for test"
}

// サービスアカウント a をリソースとして捉えて、b からアクセスできるようにする
resource "google_service_account_iam_member" "sample1" {
  service_account_id = google_service_account.a.id
  role               = "roles/iam.serviceAccountUser"
  member             = "serviceAccount:${google_service_account.b.email}"
}

// 指定したプロジェクト全体に対して、基本ロールの編集者権限をサービスアカウント a に与える
resource "google_project_iam_member" "sample2" {
  project = google_project.my_project.project_id
  role    = "roles/editor"
  member  = "serviceAccount:${google_service_account.a.email}"
}

// 指定したプロジェクト対して、cloudStorage の編集者権限をサービスアカウント c に与える
// 個々のバケットに対してのみ許可したければ専用のリソースを利用する
resource "google_project_iam_member" "sample3" {
  project = google_project.my_project.project_id
  role    = "roles/storage.admin"
  member  = "serviceAccount:${google_service_account.c.email}"
}
// c に対してもう１つロールを与える
resource "google_project_iam_member" "sample4" {
  project = google_project.my_project.project_id
  role    = "roles/run.admin"
  member  = "serviceAccount:${google_service_account.c.email}"
}
