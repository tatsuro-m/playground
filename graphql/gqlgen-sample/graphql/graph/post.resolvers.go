package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"graphql/ginctx"
	"graphql/graph/gqlmodel"
	"graphql/modelconv"
	"graphql/models"
	"graphql/service/post"
	"graphql/service/user"
	"net/http"
	"strconv"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input *gqlmodel.NewPost) (*gqlmodel.Post, error) {
	u, err := ginctx.GetUserFromGinCtx(ctx)
	if err != nil {
		return nil, errors.New(strconv.Itoa(http.StatusUnauthorized))
	}

	dbPost := models.Post{Title: input.Title, UserID: u.ID}
	p, err := post.Service{}.CreatePost(dbPost)

	return modelconv.ModelToGqlPost(&p), err
}

func (r *mutationResolver) DeletePost(ctx context.Context, input *gqlmodel.DeletePost) (string, error) {
	if _, err := ginctx.GetUserFromGinCtx(ctx); err != nil {
		return "", errors.New(strconv.Itoa(http.StatusUnauthorized))
	}

	id, _ := strconv.Atoi(input.ID)
	p, err := post.Service{}.DeleteByID(id)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(p.ID), nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*gqlmodel.Post, error) {
	dbPosts, err := post.Service{}.GetAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res := make([]*gqlmodel.Post, 0)
	for _, dp := range dbPosts {
		graphPost := modelconv.ModelToGqlPost(dp)
		owner, _ := user.Service{}.GetUserByID(dp.UserID)
		graphPost.User = modelconv.ModelToGqlUser(owner)
		res = append(res, graphPost)
	}

	return res, nil
}
