package resolver

import (
	"graphql/service/post"
	"graphql/thelper"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/99designs/gqlgen/client"

	"github.com/sebdah/goldie/v2"
)

func TestQueryResolver_Posts(t *testing.T) {
	c := createGqlClient(t)
	g := goldie.New(t)

	table := []struct {
		name  string
		query string
	}{
		{
			name:  "id と title が返ってくること",
			query: `query posts {  posts {  id  title  }}`,
		},
		{
			name:  "user も含めて返ってくること",
			query: `query users {  posts {  id title user { name email }  }}`,
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			u := thelper.InsertUser(t, 1)[0]
			thelper.InsertPost(t, 5, u.ID)

			var resp interface{}
			c.MustPost(td.query, &resp)
			g.AssertJson(t, t.Name(), resp)
		})
	}
}

func TestQueryResolver_Post(t *testing.T) {
	c := createGqlClient(t)
	g := goldie.New(t)

	table := []struct {
		name  string
		query string
	}{
		{
			name:  "指定した id の post １つだけが返ってくること",
			query: "query post {\n  post(id: 1) {\n    id\n    title\n}\n}",
		},
		{
			name:  "id 5 の post 情報が返ってくること",
			query: "query post {\n  post(id: 1) {\n    id\n    title\n    user {\n      name\n      email\n    }\n  }\n}",
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			u := thelper.InsertUser(t, 1)[0]
			thelper.InsertPost(t, 5, u.ID)

			var resp interface{}
			c.MustPost(td.query, &resp)
			g.AssertJson(t, t.Name(), resp)
		})
	}
}

func TestMutationResolver_CreatePost(t *testing.T) {
	c := createGqlClient(t)
	g := goldie.New(t)

	table := []struct {
		name  string
		query string
		input map[string]string
	}{
		{
			name:  "post が作成されること",
			query: "mutation createPost($title: String!){\n  createPost(input: {title: $title}){\n    id\n    title\n    \n    user {\n      name\n      email\n    }\n  }\n}",
			input: map[string]string{"title": "create post mutation test"},
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			var resp interface{}
			title := td.input["title"]
			c.MustPost(td.query, &resp, client.Var("title", title), thelper.AddContext(t))
			g.AssertJson(t, t.Name(), resp)

			p, _ := post.Service{}.GetByTitle(title)
			assert.Equal(t, title, p.Title)
		})
	}
}

func TestMutationResolver_AddTag(t *testing.T) {
	c := createGqlClient(t)
	g := goldie.New(t)

	table := []struct {
		name  string
		query string
		input map[string]int
	}{
		{
			name:  "post が返ってくること",
			query: "mutation addTag($post_id: ID!, $tag_id: ID!){\n  addTag(input: {post_id: $post_id, tag_id: $tag_id}){\n    id\n    title\n  }\n}",
			input: map[string]int{"post_id": 1, "tag_id": 2},
		},
		{
			name:  "user も含めて返ってくること",
			query: "mutation addTag($post_id: ID!, $tag_id: ID!){\n  addTag(input: {post_id: $post_id, tag_id: $tag_id}){\n    id\n    title\n   user{\n      name\n      email\n    }  }\n}",
			input: map[string]int{"post_id": 1, "tag_id": 2},
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			u := thelper.InsertUser(t, 1)[0]
			p := thelper.InsertPost(t, 3, u.ID)[0]
			tag := thelper.InsertTag(t, 3)[0]

			var resp interface{}
			c.MustPost(td.query, &resp,
				client.Var("post_id", p.ID),
				client.Var("tag_id", tag.ID),
				thelper.AddContext(t))

			g.AssertJson(t, t.Name(), resp)
		})
	}
}
