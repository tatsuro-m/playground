package thelper

import (
	"fmt"
	"log"
	"sqlboiler-tutorial/db"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

func FinalizeTest(t *testing.T) {
	t.Helper()

	d := db.GetDB()

	rows, err := queries.Raw("select tablename \n  from pg_tables \n  where schemaname not like 'pg_%' and schemaname != 'information_schema'").Query(d)
	if err != nil {
		fmt.Println(err)
		return
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

	excludeTables := append(make([]string, 0), "schema_migrations")

	for _, name := range tableNames {
		if notContains(excludeTables, name) {
			queries.Raw(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", name)).Exec(d)
		}
	}

	db.Close()
}

func notContains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return false
		}
	}
	return true
}
