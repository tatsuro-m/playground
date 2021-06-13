package db

import (
	"gin-gorm-tutorial/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	dsn := "host=db user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//autoMigration()
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

func autoMigration() {
	_ = db.AutoMigrate(&entity.User{})
}
