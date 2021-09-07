package post

import (
	"context"
	"sqlboiler-tutorial-mysql/db"
	"sqlboiler-tutorial-mysql/thelper"
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

func TestGetUser(t *testing.T) {
	thelper.SetupTest(t)
	defer thelper.FinalizeTest(t)

	u := thelper.InsertUser(t, 1)[0]
	p := thelper.InsertPost(t, 1, u.ID)[0]

	actual, _ := p.User().One(context.Background(), db.GetDB())
	assert.Equal(t, u, actual)
}
