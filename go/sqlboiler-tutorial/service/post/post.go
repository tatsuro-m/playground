package post

import (
	"context"
	"sqlboiler-tutorial/db"
	"sqlboiler-tutorial/models"
)

func GetAllPosts() ([]*models.Post, error) {
	return models.Posts().All(context.Background(), db.GetDB())
}
