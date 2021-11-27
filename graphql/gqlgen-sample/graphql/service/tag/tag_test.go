package tag

import (
	"fmt"
	"graphql/models"
	"graphql/service/post"
	"graphql/thelper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Posts(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		args  args
		input map[string]int
	}{
		{
			name:  "tag に紐づく posts が全件返ってくること",
			args:  struct{ id int }{id: 1},
			input: map[string]int{"postNum": 3},
		},
		{
			name:  "25 件入れても全て返ってくること",
			args:  struct{ id int }{id: 1},
			input: map[string]int{"postNum": 25},
		},
		{
			name:  "tag id に 2 を指定しても posts が全件返ってくること",
			args:  struct{ id int }{id: 2},
			input: map[string]int{"postNum": 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			u := thelper.InsertUser(t, 1)[0]
			tags := thelper.InsertTag(t, 2)
			tag := tags[0]
			anotherTag := tags[len(tags)-1]
			postNum := tt.input["postNum"]
			posts := thelper.InsertPost(t, postNum, u.ID)
			for _, p := range posts {
				s := post.Service{}

				pt := &models.PostTag{PostID: p.ID, TagID: tag.ID}
				s.AddTag(pt)
				pt = &models.PostTag{PostID: p.ID, TagID: anotherTag.ID}
				s.AddTag(pt)
			}

			got, err := Service{}.Posts(tt.args.id)
			if err != nil {
				fmt.Println(err)
			}

			assert.Len(t, got, postNum)
			for _, p := range got {
				assert.Contains(t, p.Title, "test")
				assert.True(t, p.ID > 0)
			}
		})
	}
}
