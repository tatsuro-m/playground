package post

import (
	"context"
	"fmt"
	"graphql/db"
	"graphql/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

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

func (s Service) ExistsByID(id int) bool {
	exists, _ := models.Posts(models.PostWhere.ID.EQ(id)).Exists(context.Background(), db.GetDB())
	return exists
}

func (s Service) AddTag(j *models.PostTag) error {
	ctx := context.Background()
	d := db.GetDB()
	return j.Insert(ctx, d, boil.Infer())
}

func (s Service) Tags(postID int) (models.TagSlice, error) {
	tags, err := models.Tags(
		qm.InnerJoin("post_tags pt on id = pt.tag_id"),
		qm.InnerJoin("posts p on pt.post_id = p.id"),
		qm.Where("p.id = ?", postID),
	).All(context.Background(), db.GetDB())

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return tags, nil
}
