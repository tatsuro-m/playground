package main

import (
	"pcode/db"
	"pcode/internal/seed"
)

func main() {
	db.Init()
	defer db.Close()

	seed.Exec()
}
