package test_helper

import (
	"gin-gorm-tutorial/db"
	"os"
	"testing"
)

func SetupTest(t *testing.T) {
	t.Helper()

	// テスト環境フラグが付いていない場合には強制的に落とす
	// 他環境の DB を絶対に操作しない為
	if os.Getenv("TEST_ENV") != "1" {
		t.Fatalf("テスト環境ではありませんでした。")
	}

	db.Init()
}
