resource "google_secret_manager_secret" "dsn" {
  secret_id = "${local.app_prefix}-dsn"

  replication {
    automatic = true
  }
}


resource "google_secret_manager_secret_version" "dsn" {
  secret = google_secret_manager_secret.dsn.id

  secret_data = "host=127.0.0.1 user=${google_sql_user.user.name} password=${var.POSTGRES_PASSWORD} dbname=${google_sql_database.database.name} port=5432 sslmode=disable TimeZone=Asia/Tokyo"
}
