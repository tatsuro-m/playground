package resolver

import (
	"graphql/thelper"
	"testing"

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
