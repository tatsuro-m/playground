package thelper

import (
	"fmt"
	"log"
	"pcode/db"

	"testing"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

func FinalizeTest(t *testing.T) {
	t.Helper()

	// 現状では players テーブルをクエリするだけ（編集しない）ので truncate はしない。
	// この先テストケースごとに変更するようであれば書き換える
	//truncateTable()
}

func truncateTable() {
	d := db.GetDB()
	tableNames := getTableNames()

	queries.Raw("SET FOREIGN_KEY_CHECKS = 0").Exec(d)
	for _, name := range tableNames {
		if notContains(getExcludeTables(), name) {
			c := fmt.Sprintf("TRUNCATE TABLE %s", name)
			queries.Raw(c).Exec(d)
		}
	}
	queries.Raw("SET FOREIGN_KEY_CHECKS = 1").Exec(d)

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
