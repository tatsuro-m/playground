package user

import (
	"context"
	"fmt"
	"graphql/db"
	"graphql/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Service struct{}

func (s Service) GetAll() (models.UserSlice, error) {
	return models.Users().All(context.Background(), db.GetDB())
}

func (s Service) ExistsByUID(uid string) bool {
	fmt.Println(uid)

	exists, _ := models.Users(models.UserWhere.UserID.EQ(uid)).Exists(context.Background(), db.GetDB())
	return exists
}

func (s Service) GetUserByID(id int) (*models.User, error) {
	return models.Users(models.UserWhere.ID.EQ(id)).One(context.Background(), db.GetDB())
}

func (s Service) GetUserByUID(uid string) (*models.User, error) {
	return models.Users(models.UserWhere.UserID.EQ(uid)).One(context.Background(), db.GetDB())
}

func (s Service) CreateUser(user models.User) (models.User, error) {
	ctx := context.Background()
	d := db.GetDB()

	err := user.Insert(ctx, d, boil.Infer())
	if err != nil {
		return models.User{}, err
	}

	u, err := s.GetUserByUID(user.UserID)
	if err != nil {
		return models.User{}, err
	}

	return *u, nil
}
