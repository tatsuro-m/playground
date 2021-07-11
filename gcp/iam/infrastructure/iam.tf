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
