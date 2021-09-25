package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql/graph/generated"
	"graphql/graph/model"
	"graphql/models"
	"graphql/service/post"
	"strconv"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*model.Post, error) {
	dbPost := models.Post{Title: input.Title}
	p, err := post.Service{}.CreatePost(dbPost)

	return &p, err
}

func (r *mutationResolver) DeletePost(ctx context.Context, input *model.DeletePost) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	dbPosts, err := post.Service{}.GetAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res := make([]*model.Post, 0)
	for _, dp := range dbPosts {
		graphPost := &model.Post{
			ID:        strconv.Itoa(dp.ID),
			Title:     dp.Title,
			CreatedAt: dp.CreatedAt,
			UpdatedAt: dp.UpdatedAt,
		}

		res = append(res, graphPost)
	}

	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) DeleteLink(ctx context.Context, input *model.DeleteLink) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
