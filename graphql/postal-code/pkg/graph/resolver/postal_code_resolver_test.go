package resolver

import (
	"github.com/99designs/gqlgen/client"
	"github.com/sebdah/goldie/v2"
	"pcode/pkg/thelper"
	"testing"
)

func TestQueryResolver_PostalCode(t *testing.T) {
	c := createGqlClient(t)
	g := goldie.New(t)

	t.Run("正常系", func(t *testing.T) {
		table := []struct {
			name    string
			address string
			query   string
		}{
			{
				name:    "1600004（新宿区四谷） が返ってくること",
				address: "東京都新宿区四谷",
				query:   "query address($address: String!) {\n  postalCode(address: $address) {\n    id\n    code\n}\n}",
			},
			{
				name:    "0640941（札幌市中央区旭ケ丘） が返ってくること",
				address: "北海道札幌市中央区旭ケ丘",
				query:   "query address($address: String!) {\n  postalCode(address: $address) {\n    id\n    code\n}\n}",
			},
			{
				name:    "全角スペースを入れても 0640941（札幌市中央区旭ケ丘） が返ってくること",
				address: "北海道札幌市　中央区旭ケ丘",
				query:   "query address($address: String!) {\n  postalCode(address: $address) {\n    id\n    code\n}\n}",
			},
			{
				name:    "複数の全角スペースを入れても 0640941（札幌市中央区旭ケ丘） が返ってくること",
				address: "北海道　札幌市　中央区　旭ケ丘",
				query:   "query address($address: String!) {\n  postalCode(address: $address) {\n    id\n    code\n}\n}",
			},
		}

		for _, td := range table {
			t.Run(td.name, func(t *testing.T) {
				thelper.SetupTest(t)
				defer thelper.FinalizeTest(t)
				thelper.InsertAddressData(t)

				var resp interface{}
				c.MustPost(td.query, &resp, client.Var("address", td.address))
				g.AssertJson(t, t.Name(), resp)
			})
		}
	})

	t.Run("異常系", func(t *testing.T) {
		table := []struct {
			name    string
			address string
			query   string
		}{
			{
				name:    "途中に半角スペースを入れるとエラーが返ってくること",
				address: "北海道札幌市 中央区旭ケ丘",
				query:   "query address($address: String!) {\n  postalCode(address: $address) {\n    id\n    code\n}\n}",
			},
			{
				name:    "存在しない住所を入れるとエラーが返ってくること",
				address: "東京都新宿区ほげほげ",
				query:   "query address($address: String!) {\n  postalCode(address: $address) {\n    id\n    code\n}\n}",
			},
		}

		for _, td := range table {
			t.Run(td.name, func(t *testing.T) {
				thelper.SetupTest(t)
				defer thelper.FinalizeTest(t)
				thelper.InsertAddressData(t)

				var resp interface{}
				err := c.Post(td.query, &resp, client.Var("address", td.address))
				g.AssertJson(t, t.Name(), resp)
				g.AssertJson(t, t.Name()+"_error", err)
			})
		}
	})
}
