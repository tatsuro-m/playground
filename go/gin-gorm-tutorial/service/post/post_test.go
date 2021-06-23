package post_test

import (
	"fmt"
	"gin-gorm-tutorial/db"
	"gin-gorm-tutorial/entity"
	"gin-gorm-tutorial/service/post"
	test_helper "gin-gorm-tutorial/test-helper"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetAllByUserID(t *testing.T) {
	test_helper.SetupTest(t)
	defer test_helper.FinalizeTest(t)
	d := db.GetDB()

	insertUser := func() entity.User {
		u := entity.User{FirstName: "first", LastName: "last"}
		d.Create(&u)
		return u
	}
	u := insertUser()
	insertPost := func() {
		for i := 0; i < 5; i++ {
			p := entity.Post{Title: "title" + strconv.Itoa(i), Content: "content" + strconv.Itoa(i), UserID: u.ID}
			d.Create(&p)
		}
	}
	insertPost()

	var s post.Service
	posts, err := s.GetAllByUserID(fmt.Sprintf("%v", u.ID))
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.Len(t, posts, 5)
	for _, p := range posts {
		assert.Equal(t, p.UserID, u.ID)
	}
}
