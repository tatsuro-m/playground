package post

import (
	"context"
	"graphql/db"
	"graphql/graph/model"
	"graphql/models"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Service struct{}

func (s Service) GetAll() (models.PostSlice, error) {
	return models.Posts().All(context.Background(), db.GetDB())
}

func (s Service) CreatePost(post models.Post) (model.Post, error) {
	ctx := context.Background()
	d := db.GetDB()
	err := post.Insert(ctx, d, boil.Infer())
	if err != nil {
		return model.Post{}, err
	}

	posts, err := models.Posts(models.PostWhere.Title.EQ(post.Title)).All(ctx, d)
	p := posts[len(posts)-1]

	return model.Post{ID: strconv.Itoa(p.ID), Title: p.Title, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}, nil
}
