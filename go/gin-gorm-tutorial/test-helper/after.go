package test_helper

import (
	"fmt"
	"gin-gorm-tutorial/db"
	"testing"
)

func FinalizeTest(t *testing.T) {
	t.Helper()

	var tableNames []string
	d := db.GetDB()
	d.Raw("select tablename \n  from pg_tables \n  where schemaname not like 'pg_%' and schemaname != 'information_schema';").Find(&tableNames)
	excludeTables := append(make([]string, 0), "schema_migrations")

	for _, name := range tableNames {
		if notContains(excludeTables, name) {
			d.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", name))
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
