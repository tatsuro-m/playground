package db

import (
	"database/sql"
	"fmt"
)

func Init() error {
	db, err := sql.Open("postgres", "dbname=fun user=abc")
	if err != nil {
		return err
	}
	fmt.Println(db)

	return nil
}
