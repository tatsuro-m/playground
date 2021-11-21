package post

import (
	"context"
	"graphql/db"
	"graphql/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Service struct{}

func (s Service) GetAll() (models.PostSlice, error) {
	return models.Posts().All(context.Background(), db.GetDB())
}

func (s Service) GetMyAllPosts(u models.User) (models.PostSlice, error) {
	return u.Posts().All(context.Background(), db.GetDB())
}

func (s Service) GetByID(id int) (*models.Post, error) {
	return models.FindPost(context.Background(), db.GetDB(), id)
}

func (s Service) GetByTitle(title string) (*models.Post, error) {
	return models.Posts(models.PostWhere.Title.EQ(title)).One(context.Background(), db.GetDB())
}

func (s Service) CreatePost(post models.Post) (models.Post, error) {
	ctx := context.Background()
	d := db.GetDB()
	err := post.Insert(ctx, d, boil.Infer())
	if err != nil {
		return models.Post{}, err
	}

	posts, err := models.Posts(models.PostWhere.Title.EQ(post.Title)).All(ctx, d)
	p := posts[len(posts)-1]

	return *p, nil
}

func (s Service) DeleteByID(id int) (*models.Post, error) {
	ctx := context.Background()
	d := db.GetDB()

	post, err := models.FindPost(ctx, d, id)
	if err != nil {
		return &models.Post{}, err
	}

	_, err = post.Delete(ctx, d)

	if err != nil {
		return &models.Post{}, err
	}

	return post, nil
}

func (s Service) CheckMyPost(postID, userID int) bool {
	p, err := s.GetByID(postID)
	if err != nil {
		return false
	}

	return p.UserID == userID
}

func (s Service) AddTag(j *models.PostTag) error {
	ctx := context.Background()
	d := db.GetDB()
	return j.Insert(ctx, d, boil.Infer())
}
