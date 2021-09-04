package post

import (
	"sqlboiler-tutorial/thelper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllPosts(t *testing.T) {
	thelper.SetupTest(t)
	defer thelper.FinalizeTest(t)

	u := thelper.InsertUser(t, 1)[0]
	thelper.InsertPost(t, 5, u.ID)

	posts, _ := GetAllPosts()
	assert.Equal(t, 5, len(posts))
}
