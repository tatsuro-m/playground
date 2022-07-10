package main

import (
	"context"
	"entqs/ent"
	"entqs/entutil"
	"entqs/model/user"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	c, err := ent.Open("mysql", getDSN())
	if err != nil {
		panic(err)
	}
	client := entutil.InitClient(c)
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	_, err = user.Create(&ent.User{Age: 10, Name: "test1"}, context.Background())
	if err != nil {
		log.Println(err)
	}
}

func getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		"3306",
		os.Getenv("DB_NAME"),
	)
}
