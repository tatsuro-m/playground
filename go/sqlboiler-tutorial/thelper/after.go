package thelper

import (
	"context"
	"fmt"
	"sqlboiler-tutorial/db"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

func FinalizeTest(t *testing.T) {
	t.Helper()

	var tableNames []string
	d := db.GetDB()
	ctx := context.Background()

	err := queries.Raw("select tablename \n  from pg_tables \n  where schemaname not like 'pg_%' and schemaname != 'information_schema'").Bind(ctx, d, &tableNames)
	if err != nil {
		fmt.Println(err)
		return
	}

	excludeTables := append(make([]string, 0), "schema_migrations")

	for _, name := range tableNames {
		if notContains(excludeTables, name) {
			d.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", name))
		}
	}
}

func notContains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return false
		}
	}
	return true
}
