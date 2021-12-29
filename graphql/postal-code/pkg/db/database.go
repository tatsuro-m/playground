package db

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"os"
	"pcode/pkg/util"

	_ "github.com/go-sql-driver/mysql"
)

var d *sql.DB

func Init() error {
	dsn := getDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	boilDebug()
	d = db
	return nil
}

func boilDebug() {
	if os.Getenv("BOIL_DEBUG") != "" {
		fmt.Println("enable boil debug mode")
		boil.DebugMode = true
	}
}

func getDSN() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if util.IsProd() {
		return fmt.Sprintf("%s:%s@unix(//cloudsql/%s)/%s?parseTime=true", dbUser, dbPassword, os.Getenv("INSTANCE_CONNECTION_NAME"), dbName)
	} else if util.IsDev() || util.IsTest() {
		return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
			dbUser, dbPassword, os.Getenv("DB_HOST"), dbName)
	}

	fmt.Println("please define valid APP_ENV")
	return ""
}

func GetDB() *sql.DB {
	return d
}

func Close() {
	err := d.Close()

	if err != nil {
		return
	}
}
