package test_helper

import (
	"testing"
	"time"
)

func TimeFormat(t *testing.T, time time.Time) string {
	t.Helper()

	// この形式にフォーマットして、json の time 型レスポンスの中に含まれているかをチェックする
	l := "2006-01-02"
	return time.Format(l)
}
