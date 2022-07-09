package main

import (
	"context"
	"entqs/ent"
	"entqs/entutil"
	"entqs/model/user"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	client := entutil.InitClient()
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	_, err := user.Create(&ent.User{Age: 10, Name: "test1"}, context.Background())
	if err != nil {
		log.Println(err)
	}
}
