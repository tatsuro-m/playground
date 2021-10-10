package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"graphql/code"
	"graphql/ginctx"
	"graphql/graph/gqlmodel"
	"graphql/modelconv"
	"graphql/models"
	"graphql/service/post"
	"graphql/service/user"
	"strconv"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input *gqlmodel.NewPost) (*gqlmodel.Post, error) {
	u, err := ginctx.GetUserFromGinCtx(ctx)
	if err != nil {
		return nil, errors.New(code.NotAuthorize)
	}

	dbPost := models.Post{Title: input.Title, UserID: u.ID}
	p, err := post.Service{}.CreatePost(dbPost)

	return modelconv.ModelToGqlPost(&p), err
}

func (r *mutationResolver) DeletePost(ctx context.Context, input *gqlmodel.DeletePost) (string, error) {
	u, err := ginctx.GetUserFromGinCtx(ctx)
	if err != nil {
		return "", errors.New(code.NotAuthorize)
	}

	var s post.Service
	id, _ := strconv.Atoi(input.ID)

	if !s.CheckMyPost(id, u.ID) {
		return "", errors.New(code.Forbid)
	}

	p, err := s.DeleteByID(id)
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

func (r *queryResolver) Post(ctx context.Context, id string) (*gqlmodel.Post, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New(code.InvalidID)
	}

	p, err := post.Service{}.GetByID(i)
	if err != nil {
		return nil, errors.New(code.ModelError)
	}

	res := modelconv.ModelToGqlPost(p)
	u, err := user.Service{}.GetUserByID(p.UserID)
	res.User = modelconv.ModelToGqlUser(u)

	return res, nil
}
