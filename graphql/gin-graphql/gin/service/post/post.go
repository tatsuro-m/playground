package post

import (
	"context"
	"gin/db"
	"gin/models"
)

type Service struct{}

func (s Service) GetAll() (models.PostSlice, error) {
	return models.Posts().All(context.Background(), db.GetDB())
}
