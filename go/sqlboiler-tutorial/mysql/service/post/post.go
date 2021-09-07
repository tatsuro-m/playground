package post

import (
	"context"
	"sqlboiler-tutorial-mysql/db"
	"sqlboiler-tutorial-mysql/models"
)

func GetAllPosts() ([]*models.Post, error) {
	return models.Posts().All(context.Background(), db.GetDB())
}

func GetUser(postID int) (*models.User, error) {
	ctx := context.Background()
	d := db.GetDB()

	p, err := models.FindPost(ctx, d, postID)
	if err != nil {
		return &models.User{}, err
	}

	return p.User().One(ctx, d)
}
