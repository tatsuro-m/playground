package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/friendsofgo/errors"
	"graphql/graph/gqlmodel"
	"graphql/service/user"
	"strconv"
)

func (r *queryResolver) Users(ctx context.Context) ([]*gqlmodel.User, error) {
	dbUsers, err := user.Service{}.GetAll()
	if err != nil {
		return nil, errors.New("model error")
	}

	gqlUsers := make([]*gqlmodel.User, 0)
	for _, du := range dbUsers {
		gu := &gqlmodel.User{ID: strconv.Itoa(du.ID), Email: du.Email, Name: du.Name, Picture: du.Picture, CreatedAt: du.CreatedAt, UpdatedAt: du.UpdatedAt}
		gqlUsers = append(gqlUsers, gu)
	}

	return gqlUsers, nil
}
