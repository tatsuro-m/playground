package main

import (
	"fmt"
	"sqlboiler-tutorial/db"
)

func main() {
	err := db.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("動作は test を使って確認！")
}
