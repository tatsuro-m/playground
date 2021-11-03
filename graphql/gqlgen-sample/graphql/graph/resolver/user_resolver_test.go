package resolver

import (
	"fmt"
	"graphql/thelper"
	"testing"
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
		fmt.Println(resp)
	})
}
