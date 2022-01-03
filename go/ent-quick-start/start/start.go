package main

import (
	"context"
	"entqs/ent"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", getDSN())
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// DDL を標準出力に吐き出してみる
	client.Schema.WriteTo(context.Background(), os.Stdout)
	// オートマイグレーションツールを実行する
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	CreateUser(context.Background(), client)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func getDSN() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		dbUser, dbPassword, os.Getenv("DB_HOST"), dbName)
}
