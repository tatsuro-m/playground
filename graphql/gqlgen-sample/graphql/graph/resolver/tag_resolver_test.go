package resolver

import (
	"graphql/models"
	"graphql/service/post"
	"graphql/thelper"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/sebdah/goldie/v2"
)

func TestQueryResolver_TagPosts(t *testing.T) {
	c := createGqlClient(t)
	g := goldie.New(t)

	table := []struct {
		name  string
		query string
		input map[string]int
	}{
		{
			name:  "tag に紐づく post が全て帰ってくること",
			query: "query tagPosts($tag_id: ID!){\n  tagPosts(tag_id: $tag_id){\n    id\n    title\n    \n    user {\n      id\n      name\n    }\n  }\n}",
			input: map[string]int{"tag_id": 1, "postNum": 3},
		},
		{
			name:  "存在しない tag_id を指定するとエラーになること",
			query: "query tagPosts($tag_id: ID!){\n  tagPosts(tag_id: $tag_id){\n    id\n    title\n    \n    user {\n      id\n      name\n    }\n  }\n}",
			input: map[string]int{"tag_id": 9999, "postNum": 3},
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			u := thelper.InsertUser(t, 1)[0]
			n := td.input["postNum"]
			posts := thelper.InsertPost(t, n, u.ID)
			tags := thelper.InsertTag(t, 2)
			tag := tags[0]
			anotherTag := tags[len(tags)-1]

			for _, p := range posts {
				s := post.Service{}
				pt := models.PostTag{PostID: p.ID, TagID: tag.ID}
				s.AddTag(&pt)
				pt = models.PostTag{PostID: p.ID, TagID: anotherTag.ID}
				s.AddTag(&pt)
			}

			var resp interface{}
			err := c.Post(td.query, &resp, client.Var("tag_id", td.input["tag_id"]))

			if err != nil {
				g.AssertJson(t, "Error_"+t.Name(), err)
			}

			g.AssertJson(t, t.Name(), resp)
		})
	}
}
