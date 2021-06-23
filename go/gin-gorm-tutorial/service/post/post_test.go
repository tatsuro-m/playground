package post_test

import (
	"fmt"
	"gin-gorm-tutorial/service/post"
	test_helper "gin-gorm-tutorial/test-helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetAllByUserID(t *testing.T) {
	test_helper.SetupTest(t)
	defer test_helper.FinalizeTest(t)

	u := test_helper.InsertUser(t, 1)[0]
	test_helper.InsertPost(t, 5, u)

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
