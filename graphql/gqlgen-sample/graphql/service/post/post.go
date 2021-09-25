package post

import (
	"context"
	"graphql/db"
	"graphql/models"
)

type Service struct{}

func (s Service) GetAll() (models.PostSlice, error) {
	return models.Posts().All(context.Background(), db.GetDB())
}
