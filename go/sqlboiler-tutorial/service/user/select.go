package user

import (
	"context"
	"sqlboiler-tutorial/db"
	"sqlboiler-tutorial/models"
)

func GetAllUsers() ([]*models.User, error) {
	return models.Users().All(context.Background(), db.GetDB())
}

func GetUserByName(name string) (*models.User, error) {
	return models.Users(models.UserWhere.Name.EQ(name)).One(context.Background(), db.GetDB())
}

func GetUserByID(id int) (*models.User, error) {
	return models.FindUser(context.Background(), db.GetDB(), id)
}

func Count() (int64, error) {
	return models.Users().Count(context.Background(), db.GetDB())
}
