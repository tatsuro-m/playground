package thelper

import (
	"pcode/pkg/db"
	"pcode/pkg/util"
	"testing"
)

func SetupTest(t *testing.T) {
	t.Helper()

	if !util.IsTest() {
		t.Fatalf("テスト環境ではありませんでした。")
	}

	db.Init()
}
