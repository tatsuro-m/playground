package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"graphql/code"
	"graphql/ginctx"
	"graphql/graph"
	"graphql/graph/gqlmodel"
	"graphql/modelconv"
	"graphql/models"
	"graphql/service/post"
	"graphql/service/user"
	"strconv"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input *gqlmodel.NewPost) (*gqlmodel.Post, error) {
	u, _ := ginctx.GetUserFromGinCtx(ctx)
	dbPost := models.Post{Title: input.Title, UserID: u.ID}
	p, err := post.Service{}.CreatePost(dbPost)

	gqlPost := modelconv.ModelToGqlPost(&p)
	gqlPost.User = modelconv.ModelToGqlUser(u)

	return gqlPost, err
}

func (r *mutationResolver) DeletePost(ctx context.Context, input *gqlmodel.DeletePost) (string, error) {
	u, _ := ginctx.GetUserFromGinCtx(ctx)
	var s post.Service
	id, _ := strconv.Atoi(input.ID)

	if !s.CheckMyPost(id, u.ID) {
		return "", graph.NewGqlError("not my post", code.AuthorizationErr)
	}

	p, err := s.DeleteByID(id)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(p.ID), nil
}

func (r *mutationResolver) AddTag(ctx context.Context, input *gqlmodel.AddTag) (*gqlmodel.Post, error) {
	pID, _ := strconv.Atoi(input.PostID)
	tID, _ := strconv.Atoi(input.TagID)
	j := models.PostTag{PostID: pID, TagID: tID}

	service := post.Service{}
	err := service.AddTag(&j)
	if err != nil {
		return nil, err
	}

	p, err := service.GetByID(pID)
	gqlPost := graph.SetUser(p)

	return gqlPost, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*gqlmodel.Post, error) {
	dbPosts, err := post.Service{}.GetAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res := make([]*gqlmodel.Post, 0)
	for _, dp := range dbPosts {
		gqlPost := graph.SetUser(dp)
		res = append(res, gqlPost)
	}

	return res, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*gqlmodel.Post, error) {
	i, err := strconv.Atoi(id)
	if err != nil {
		return nil, graph.NewGqlError(err.Error(), code.InternalErr)
	}

	p, err := post.Service{}.GetByID(i)
	if err != nil {
		return nil, graph.NewGqlError(err.Error(), code.InternalErr)
	}

	res := modelconv.ModelToGqlPost(p)
	u, err := user.Service{}.GetUserByID(p.UserID)
	res.User = modelconv.ModelToGqlUser(u)

	return res, nil
}

func (r *queryResolver) Tags(ctx context.Context, input *gqlmodel.Tags) ([]*gqlmodel.Tag, error) {
	id, err := strconv.Atoi(input.PostID)
	if err != nil {
		return nil, err
	}

	s := post.Service{}
	if !s.ExistsByID(id) {
		return nil, errors.New(fmt.Sprintf("record (post_id = %d) not exists", id))
	}
	tags, err := s.Tags(id)
	if err != nil {
		return nil, err
	}

	gqlTags := make([]*gqlmodel.Tag, 0)
	for _, tag := range tags {
		gqlTags = append(gqlTags, modelconv.ModelToGqlTag(tag))
	}

	return gqlTags, nil
}
