package post

import (
	"github.com/stretchr/testify/assert"
	"graphql/thelper"
	"testing"
)

func TestService_GetAll(t *testing.T) {
	table := []struct {
		name      string
		insertNum int
	}{
		{
			name:      "２件入れたら２件返ってくること",
			insertNum: 2,
		},
		{
			name:      "100件入れても全て返ってくること",
			insertNum: 100,
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			u := thelper.InsertUser(t, 1)[0]
			thelper.InsertPost(t, td.insertNum, u.ID)

			posts, _ := Service{}.GetAll()
			assert.Equal(t, td.insertNum, len(posts))
		})
	}
}
