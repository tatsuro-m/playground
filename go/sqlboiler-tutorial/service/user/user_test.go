package user

import (
	"sqlboiler-tutorial/thelper"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	thelper.SetupTest(t)
	defer thelper.FinalizeTest(t)

}
