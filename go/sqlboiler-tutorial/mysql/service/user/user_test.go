package user

import (
	"sqlboiler-tutorial-mysql/thelper"
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

func TestGetPostsByUserID(t *testing.T) {
	thelper.SetupTest(t)
	defer thelper.FinalizeTest(t)

	users := thelper.InsertUser(t, 2)
	u := users[0]
	thelper.InsertPost(t, 5, u.ID)
	otherUser := users[len(users)-1]
	thelper.InsertPost(t, 2, otherUser.ID)

	// それぞれの user の posts だけ返ってくるか確認する
	posts, _ := GetPostsByUserID(u.ID)
	assert.Equal(t, 5, len(posts))

	posts, _ = GetPostsByUserID(otherUser.ID)
	assert.Equal(t, 2, len(posts))
}
