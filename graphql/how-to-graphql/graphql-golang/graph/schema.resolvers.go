package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tatsuro-m/hackernews/graph/generated"
	"github.com/tatsuro-m/hackernews/graph/model"
	"github.com/tatsuro-m/hackernews/internal/auth"
	"github.com/tatsuro-m/hackernews/internal/links"
	"github.com/tatsuro-m/hackernews/internal/users"
	"github.com/tatsuro-m/hackernews/pkg/jwt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input *model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	var link links.Link
	link.Title = input.Title
	link.Address = input.Address
	link.User = user
	linkID := link.Save()

	graphqlUser := &model.User{
		ID:   user.ID,
		Name: user.Username,
	}

	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title: link.Title, Address: link.Address, User: graphqlUser}, nil
}

func (r *mutationResolver) DeleteLink(ctx context.Context, input *model.DeleteLink) (bool, error) {
	success := links.DeleteByID(input.ID)
	return success, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	correct := user.Authenticate()

	if !correct {
		return "", &users.WrongUsernameOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()

	for _, link := range dbLinks {
		graphqlUser := &model.User{ID: link.User.ID, Name: link.User.Username}

		resultLinks = append(resultLinks, &model.Link{
			ID:      link.ID,
			Title:   link.Title,
			Address: link.Address,
			User:    graphqlUser,
		})
	}

	return resultLinks, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var resultUsers []*model.User
	var dbUsers []users.User
	dbUsers = users.GetAll()

	for _, user := range dbUsers {
		resultUsers = append(resultUsers, &model.User{ID: user.ID, Name: user.Username})
	}

	return resultUsers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
