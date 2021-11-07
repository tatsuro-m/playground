package resolver

import (
	"graphql/thelper"
	"testing"

	"github.com/sebdah/goldie/v2"
)

func TestQueryResolver_Users(t *testing.T) {
	c := createGqlClient(t)
	g := goldie.New(t)

	table := []struct {
		name  string
		query string
	}{
		{
			name:  "name と picture が返ってくること",
			query: `query users {  users {  name  picture  }}`,
		},
		{
			name:  "email と picture だけが返ってくること",
			query: `query users {  users {  email  picture  }}`,
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			thelper.InsertUser(t, 5)
			var resp interface{}

			c.MustPost(td.query, &resp)
			g.AssertJson(t, t.Name(), resp)
		})
	}
}
