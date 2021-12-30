package resolver

import (
	"github.com/99designs/gqlgen/client"
	"github.com/sebdah/goldie/v2"
	"pcode/pkg/thelper"
	"testing"
)

func TestQueryResolver_Address(t *testing.T) {
	c := createGqlClient(t)
	g := goldie.New(t)

	table := []struct {
		name       string
		postalCode int
		query      string
	}{
		{
			name:       "新宿区四谷が返ってくること",
			postalCode: 1600004,
			query:      "query address($postal_code: Int!) {\n  address(postal_code: $postal_code) {\n    id\n    name\n}\n}",
		},
	}

	for _, td := range table {
		t.Run(td.name, func(t *testing.T) {
			thelper.SetupTest(t)
			defer thelper.FinalizeTest(t)

			var resp interface{}
			c.MustPost(td.query, &resp, client.Var("postal_code", td.postalCode))
			g.AssertJson(t, t.Name(), resp)
		})
	}
}
