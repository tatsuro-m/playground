package thelper

import (
	"os"
	"sqlboiler-tutorial/db"
	"testing"
)

func SetupTest(t *testing.T) {
	t.Helper()

	if os.Getenv("TEST_ENV") != "1" {
		t.Fatalf("テスト環境ではありませんでした。")
	}

	db.Init()
}
