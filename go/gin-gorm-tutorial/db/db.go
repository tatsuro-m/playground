package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open(postgres.Open(os.Getenv("DEV_DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	db, err := db.DB()
	if err != nil {
		return
	}

	if err := db.Close(); err != nil {
		return
	}
}
