package db

import (
	"database/sql"
	"fmt"
	"os"
)

var d *sql.DB

func Init() error {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	d = db
	return nil
}

func GetDB() *sql.DB {
	return d
}
