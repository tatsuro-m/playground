package seed

import (
	"pcode/thelper"
	"testing"
)

func TestExec(t *testing.T) {
	t.Run("seed が登録されること", func(t *testing.T) {
		thelper.SetupTest(t)
		defer thelper.FinalizeTest(t)

		Exec()
	})
}
