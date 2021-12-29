package thelper

import (
	"pcode/db"
	"pcode/util"
	"testing"
)

func SetupTest(t *testing.T) {
	t.Helper()

	if !util.IsTest() {
		t.Fatalf("テスト環境ではありませんでした。")
	}

	db.Init()
}
