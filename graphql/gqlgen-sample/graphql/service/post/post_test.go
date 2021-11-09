package post

import (
	"graphql/thelper"
	"testing"

	"github.com/stretchr/testify/assert"
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

	t.Run("id 順にソートされていること", func(t *testing.T) {
		thelper.SetupTest(t)
		defer thelper.FinalizeTest(t)

		u := thelper.InsertUser(t, 1)[0]
		num := 10
		thelper.InsertPost(t, num, u.ID)

		var expected []int
		for i := 0; i < 10; i++ {
			expected = append(expected, i+1)
		}

		posts, _ := Service{}.GetAll()
		var actual []int
		for _, p := range posts {
			actual = append(actual, p.ID)
		}

		for i, id := range expected {
			assert.Equal(t, id, actual[i])
		}
	})
}

func TestService_GetMyAllPosts(t *testing.T) {
	table := []struct {
		name      string
		insertNum int
	}{
		{
			name:      "２件入れたら２件返ってくること",
			insertNum: 2,
		},
		{
			name:      "30件入れても全て返ってくること",
			insertNum: 30,
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			users := thelper.InsertUser(t, 2)
			owner := users[0]
			anotherU := users[1]

			thelper.InsertPost(t, td.insertNum, owner.ID)
			thelper.InsertPost(t, 3, anotherU.ID)

			posts, _ := Service{}.GetMyAllPosts(owner)
			assert.Equal(t, td.insertNum, len(posts))
		})
	}
}
