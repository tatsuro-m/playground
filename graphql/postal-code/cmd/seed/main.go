package main

import (
	"fmt"
	"os"
	"pcode/internal/seed"
	"pcode/pkg/db"
)

func main() {
	db.Init()
	defer db.Close()

	err := seed.Exec()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
