package thelper

import (
	"fmt"
	"log"
	"sqlboiler-tutorial-mysql/db"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

func FinalizeTest(t *testing.T) {
	t.Helper()

	truncateTable()
}

func truncateTable() {
	d := db.GetDB()
	tableNames := getTableNames()

	for _, name := range tableNames {
		if notContains(getExcludeTables(), name) {
			cmds := []string{"SET FOREIGN_KEY_CHECKS = 0", fmt.Sprintf("TRUNCATE TABLE %s CASCADE", name), "SET FOREIGN_KEY_CHECKS = 1"}
			for _, c := range cmds {
				queries.Raw(c).Exec(d)
			}
		}
	}

	db.Close()
}

func getExcludeTables() []string {
	return []string{"schema_migrations"}
}

func getTableNames() []string {
	d := db.GetDB()

	rows, err := queries.Raw("SHOW TABLES;").Query(d)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	var tableNames []string
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}

		tableNames = append(tableNames, tableName)
	}

	return tableNames
}

func notContains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return false
		}
	}
	return true
}
