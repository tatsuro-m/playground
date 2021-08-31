package main

import (
	"fmt"
	"sqlboiler-tutorial/db"
	"sqlboiler-tutorial/models"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(models.UserColumns.ID)
	db.Init()
}
