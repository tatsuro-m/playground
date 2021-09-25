package main

import (
	"fmt"
	"gin/db"
)

func main() {
	err := db.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Hello World!")
}
