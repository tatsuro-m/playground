package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql/graph/gqlmodel"
)

func (r *queryResolver) TagPosts(ctx context.Context, tagID string) ([]*gqlmodel.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
