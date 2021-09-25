package main

import (
	"fmt"
	"gin/db"
	"gin/server"
)

func main() {
	err := db.Init()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	server.Init()
}
