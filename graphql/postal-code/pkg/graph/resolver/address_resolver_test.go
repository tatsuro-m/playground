package resolver

import (
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/sebdah/goldie/v2"
	"pcode/pkg/thelper"
	"testing"
)

func TestQueryResolver_Address(t *testing.T) {
	c := createGqlClient(t)
	g := goldie.New(t)

	t.Run("正常系", func(t *testing.T) {
		table := []struct {
			name       string
			postalCode string
			query      string
		}{
			{
				name:       "新宿区四谷 が返ってくること",
				postalCode: "1600004",
				query:      "query address($postal_code: String!) {\n  address(postal_code: $postal_code) {\n    id\n    name\n}\n}",
			},
			{
				name:       "札幌市中央区旭ケ丘 が返ってくること",
				postalCode: "0640941",
				query:      "query address($postal_code: String!) {\n  address(postal_code: $postal_code) {\n    id\n    name\n}\n}",
			},
		}

		for _, td := range table {
			t.Run(td.name, func(t *testing.T) {
				thelper.SetupTest(t)
				defer thelper.FinalizeTest(t)
				thelper.InsertAddressData(t)

				var resp interface{}
				c.MustPost(td.query, &resp, client.Var("postal_code", td.postalCode))
				g.AssertJson(t, t.Name(), resp)
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		table := []struct {
			name       string
			postalCode string
			query      string
		}{
			{
				name:       "１0桁を指定するとエラーになること",
				postalCode: "0000000000",
				query:      "query address($postal_code: String!) {\n  address(postal_code: $postal_code) {\n    id\n    name\n}\n}",
			},
			{
				name:       "存在しない郵便番号を指定するとエラーになること",
				postalCode: "1111111",
				query:      "query address($postal_code: String!) {\n  address(postal_code: $postal_code) {\n    id\n    name\n}\n}",
			},
			{
				name:       "アルファベットを指定するとエラーになること",
				postalCode: "aaaaaaa",
				query:      "query address($postal_code: String!) {\n  address(postal_code: $postal_code) {\n    id\n    name\n}\n}",
			},
			{
				name:       "ひらがなを指定するとエラーになること",
				postalCode: "あああああああ",
				query:      "query address($postal_code: String!) {\n  address(postal_code: $postal_code) {\n    id\n    name\n}\n}",
			},
		}

		for _, td := range table {
			t.Run(td.name, func(t *testing.T) {
				thelper.SetupTest(t)
				defer thelper.FinalizeTest(t)
				thelper.InsertAddressData(t)

				var resp interface{}
				err := c.Post(td.query, &resp, client.Var("postal_code", td.postalCode))
				fmt.Println(t.Name())
				g.AssertJson(t, t.Name(), resp)
				g.AssertJson(t, t.Name()+"_error", err)
			})
		}
	})
}
