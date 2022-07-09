package entutil

import (
	"entqs/ent"
	"fmt"
	"os"
)

var c *ent.Client

func InitClient() *ent.Client {
	client, err := ent.Open("mysql", getDSN())
	if err != nil {
		panic(err)
	}

	c = client
	return c
}

func GetEntClient() *ent.Client {
	return c
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
