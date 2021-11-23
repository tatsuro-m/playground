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
		name         string
		query        string
		input        map[string]string
		setValidUser bool
	}{
		{
			name:         "post が作成されること",
			query:        "mutation createPost($title: String!){\n  createPost(input: {title: $title}){\n    id\n    title\n    \n    user {\n      name\n      email\n    }\n  }\n}",
			input:        map[string]string{"title": "create post mutation test"},
			setValidUser: true,
		},
		{
			name:         "認証が必須であること",
			query:        "mutation createPost($title: String!){\n  createPost(input: {title: $title}){\n    id\n    title\n    \n    user {\n      name\n      email\n    }\n  }\n}",
			input:        map[string]string{"title": "create post mutation test"},
			setValidUser: false,
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			u := thelper.InsertUser(t, 1)[0]
			var op client.Option
			if td.setValidUser {
				op = thelper.SetUserToContext(t, &u)
			} else {
				op = thelper.SetEmptyUserToContext(t)
			}

			var resp interface{}
			title := td.input["title"]
			err := c.Post(td.query, &resp, client.Var("title", title), op)
			if err != nil {
				g.AssertJson(t, "Error"+t.Name(), err)
			} else {
				p, _ := post.Service{}.GetByTitle(title)
				assert.Equal(t, title, p.Title)
			}

			g.AssertJson(t, t.Name(), resp)
		})
	}
}

func TestMutationResolver_DeletePost(t *testing.T) {
	c := createGqlClient(t)
	g := goldie.New(t)

	table := []struct {
		name          string
		query         string
		input         map[string]int
		authenticated bool
		myPost        bool
	}{
		{
			name:          "指定した id の post が削除できること",
			query:         "mutation deletePost($post_id: ID!){\n  deletePost(input: {id: $post_id})\n}",
			input:         map[string]int{"post_id": 1},
			authenticated: true,
			myPost:        true,
		},
		{
			name:          "未認証だと post が削除できないこと",
			query:         "mutation deletePost($post_id: ID!){\n  deletePost(input: {id: $post_id})\n}",
			input:         map[string]int{"post_id": 1},
			authenticated: false,
			myPost:        true,
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			users := thelper.InsertUser(t, 2)
			owner := users[0]
			anotherUser := users[1]
			thelper.InsertPost(t, 1, owner.ID)
			thelper.InsertPost(t, 1, anotherUser.ID)

			var op client.Option
			if td.authenticated && td.myPost {
				op = thelper.SetUserToContext(t, &owner)
			} else if td.authenticated {
				op = thelper.SetUserToContext(t, &anotherUser)
			} else {
				op = thelper.SetEmptyUserToContext(t)
			}

			var resp interface{}
			pID := td.input["post_id"]
			err := c.Post(td.query, &resp,
				client.Var("post_id", pID),
				op,
			)

			if err != nil {
				g.AssertJson(t, "Error_"+t.Name(), err)
				assert.Equal(t, true, post.Service{}.ExistsByID(pID))
			} else {
				assert.Equal(t, false, post.Service{}.ExistsByID(pID))
			}

			g.AssertJson(t, t.Name(), resp)
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
		{
			name:  "存在しない post_id を指定するとエラーになること",
			query: "mutation addTag($post_id: ID!, $tag_id: ID!){\n  addTag(input: {post_id: $post_id, tag_id: $tag_id}){\n    id\n    title\n   user{\n      name\n      email\n    }  }\n}",
			input: map[string]int{"post_id": 9999, "tag_id": 2},
		},
		{
			name:  "存在しない tag_id を指定するとエラーになること",
			query: "mutation addTag($post_id: ID!, $tag_id: ID!){\n  addTag(input: {post_id: $post_id, tag_id: $tag_id}){\n    id\n    title\n   user{\n      name\n      email\n    }  }\n}",
			input: map[string]int{"post_id": 1, "tag_id": 9999},
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			u := thelper.InsertUser(t, 1)[0]
			thelper.InsertPost(t, 3, u.ID)
			thelper.InsertTag(t, 3)

			var resp interface{}
			err := c.Post(td.query, &resp,
				client.Var("post_id", td.input["post_id"]),
				client.Var("tag_id", td.input["tag_id"]),
			)

			if err != nil {
				// err が存在するなら res body に合わせてエラーの中身もテストしたいので。
				g.AssertJson(t, "Error_"+t.Name(), err)
			}

			g.AssertJson(t, t.Name(), resp)
		})
	}
}
