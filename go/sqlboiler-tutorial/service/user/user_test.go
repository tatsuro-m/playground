package user

import (
	"sqlboiler-tutorial/thelper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	thelper.SetupTest(t)
	defer thelper.FinalizeTest(t)

	recordNum := 5
	thelper.InsertUser(t, recordNum)
	users, _ := GetAllUsers()
	actual := len(users)

	assert.Equal(t, recordNum, actual)
}
