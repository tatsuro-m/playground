package thelper

import (
	"fmt"
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
	tableNames := GetTableNames()

	for _, name := range tableNames {
		c := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", name)
		queries.Raw(c).Exec(d)
	}

	db.Close()
}

func GetTableNames() []string {
	// 現状ホワイトリスト形式で対象のテーブルを列挙しているが、自動的に取得できるようにしたい。。。
	// mysql だと最初から作成されているマスター的なテーブルがあるからそこも削除されないように。
	tableNames := []string{"users", "posts"}
	return tableNames
}
