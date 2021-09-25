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

	// 一意性制約の付いていない title カラムで検索して返しているだけなので今回入れたレコードか保証できないのが微妙。
	p, err := models.Posts(models.PostWhere.Title.EQ(post.Title)).One(ctx, d)
	return model.Post{ID: strconv.Itoa(p.ID), Title: p.Title, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}, nil
}
