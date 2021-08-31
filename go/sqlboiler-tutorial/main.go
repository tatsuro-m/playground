package main

import (
	"fmt"
	"sqlboiler-tutorial/db"
	"sqlboiler-tutorial/service/user"
)

func main() {
	err := db.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	d := db.GetDB()

	user.Insert(d)
}
