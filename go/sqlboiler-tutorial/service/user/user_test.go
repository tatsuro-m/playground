package user

import (
	"context"
	"math"
	"sqlboiler-tutorial/db"
	"sqlboiler-tutorial/models"
	"sqlboiler-tutorial/thelper"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	thelper.SetupTest(t)
	defer thelper.FinalizeTest(t)

	recordNum := 5
	thelper.InsertUser(t, recordNum)
	actual, _ := models.Users().Count(context.Background(), db.GetDB())

	if recordNum != Int64ToInt(actual) {
		t.Error("error!")
	}
}

func Int64ToInt(i int64) int {
	if i < math.MinInt32 || i > math.MaxInt32 {
		return 0
	} else {
		return int(i)
	}
}
