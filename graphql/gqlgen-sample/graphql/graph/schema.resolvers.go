package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql/graph/generated"
	"graphql/graph/model"
	"graphql/internal/url"
	"io/ioutil"
	"net/http"
)

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	req, err := http.NewRequest(http.MethodGet, url.GetAPIPath("/posts"), nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("手に入れた body")
	fmt.Println(body)

	return nil, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
