package main

import (
	"pcode/internal/seed"
	"pcode/pkg/db"
)

func main() {
	db.Init()
	defer db.Close()

	seed.Exec()
}
