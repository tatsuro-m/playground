package post

import (
	"fmt"
	"graphql/models"
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

func TestService_Tags(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		args  args
		input map[string]int
	}{
		{
			name:  "post に紐づく tags が全件帰ってくること",
			args:  struct{ id int }{id: 1},
			input: map[string]int{"tagNum": 3},
		},
		{
			name:  "25 件入れても全て返ってくること",
			args:  struct{ id int }{id: 1},
			input: map[string]int{"tagNum": 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			u := thelper.InsertUser(t, 1)[0]
			posts := thelper.InsertPost(t, 2, u.ID)
			p := posts[0]
			anotherP := posts[len(posts)-1]

			tagNum := tt.input["tagNum"]
			tags := thelper.InsertTag(t, tagNum)
			s := Service{}
			for _, tag := range tags {
				pt := &models.PostTag{PostID: p.ID, TagID: tag.ID}
				s.AddTag(pt)

				// tag を他の post とも関連付ける
				pt = &models.PostTag{PostID: anotherP.ID, TagID: tag.ID}
				s.AddTag(pt)
			}

			got, err := s.Tags(tt.args.id)
			if err != nil {
				fmt.Println(err)
			}

			assert.Len(t, got, tagNum)
			for _, tag := range got {
				assert.Contains(t, tag.Name, "test")
				assert.IsType(t, 0, tag.ID)
			}
		})
	}
}
