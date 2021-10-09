package db

import (
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"os"

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
	} else {
		fmt.Println("disable boil debug mode")
	}
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
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
