package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql/graph/generated"
	"graphql/graph/model"
	"graphql/middleware"
	"graphql/models"
	"graphql/service/post"
	"net/http"
	"strconv"

	"github.com/friendsofgo/errors"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*model.Post, error) {
	user := middleware.ForContext(ctx)
	fmt.Println(user)

	if user == nil {
		fmt.Println("user が設定されていないようです")
		return nil, errors.New(strconv.Itoa(http.StatusNotFound))
	} else {
		fmt.Println("user がセットされていたので表示してみる")
		fmt.Println(user)
	}

	dbPost := models.Post{Title: input.Title}
	p, err := post.Service{}.CreatePost(dbPost)

	return &model.Post{ID: strconv.Itoa(p.ID), Title: p.Title, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}, err
}

func (r *mutationResolver) DeletePost(ctx context.Context, input *model.DeletePost) (string, error) {
	id, _ := strconv.Atoi(input.ID)
	p, err := post.Service{}.DeleteByID(id)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(p.ID), nil
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
