package resolver

import (
	"fmt"
	"graphql/thelper"
	"testing"

	"github.com/sebdah/goldie/v2"
)

func TestQueryResolver_Users(t *testing.T) {
	c := createGqlClient(t)

	t.Run("全ての user が返ってくること", func(t *testing.T) {
		thelper.SetupTest(t)
		defer thelper.FinalizeTest(t)

		thelper.InsertUser(t, 5)
		var resp interface{}
		q := `
query users {
  users {
    name
    picture
  }
}`

		c.MustPost(q, &resp)

		g := goldie.New(t)
		g.AssertJson(t, "example", resp)
		fmt.Println(resp)
	})
}
