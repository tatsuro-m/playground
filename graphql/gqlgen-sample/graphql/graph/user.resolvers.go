package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/friendsofgo/errors"
	"graphql/graph/gqlmodel"
	"graphql/modelconv"
	"graphql/service/user"
)

func (r *queryResolver) Users(ctx context.Context) ([]*gqlmodel.User, error) {
	dbUsers, err := user.Service{}.GetAll()
	if err != nil {
		return nil, errors.New("model error")
	}

	gqlUsers := make([]*gqlmodel.User, 0)
	for _, du := range dbUsers {
		gu := modelconv.ModelToGqlUser(du)
		gqlUsers = append(gqlUsers, gu)
	}

	return gqlUsers, nil
}
