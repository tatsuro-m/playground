package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql/code"
	"graphql/graph"
	"graphql/graph/gqlmodel"
	"graphql/modelconv"
	"graphql/service/user"
)

func (r *queryResolver) Users(ctx context.Context) ([]*gqlmodel.User, error) {
	modelUsers, err := user.Service{}.GetAll()
	if err != nil {
		return nil, graph.NewGqlError(err.Error(), code.InternalErr)
	}

	gqlUsers := make([]*gqlmodel.User, 0)
	for _, mu := range modelUsers {
		gu := modelconv.ModelToGqlUser(mu)
		gqlUsers = append(gqlUsers, gu)
	}

	return gqlUsers, nil
}
