package user

import (
	"context"
	"sqlboiler-tutorial/db"
	"sqlboiler-tutorial/models"
)

func DeleteByID(id int) (*models.User, error) {
	c := context.Background()
	u, err := models.FindUser(c, db.GetDB(), id)
	if err != nil {
		return nil, err
	}

	_, err = u.Delete(c, db.GetDB())
	if err != nil {
		return nil, err
	}

	return u, nil
}
