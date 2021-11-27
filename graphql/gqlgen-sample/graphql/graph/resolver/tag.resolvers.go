package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql/graph"
	"graphql/graph/gqlmodel"
	"graphql/service/tag"
	"strconv"
)

func (r *queryResolver) TagPosts(ctx context.Context, tagID string) ([]*gqlmodel.Post, error) {
	tID, err := strconv.Atoi(tagID)
	if err != nil {
		return nil, err
	}

	mPosts, err := tag.Service{}.Posts(tID)
	if err != nil {
		return nil, err
	}

	gqlPosts := make([]*gqlmodel.Post, 0)
	for _, mp := range mPosts {
		gqlPosts = append(gqlPosts, graph.SetUser(mp))
	}

	return gqlPosts, nil
}
