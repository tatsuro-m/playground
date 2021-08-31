package user

import (
	"context"
	"sqlboiler-tutorial/db"
	"sqlboiler-tutorial/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func UpdateByID(id int) (*models.User, error) {
	c := context.Background()
	u, err := models.FindUser(c, db.GetDB(), id)
	if err != nil {
		return nil, err
	}

	//	実際は引数で更新する User 構造体を取得したりするが、今回は固定値
	u.Name = "updated name!"
	_, err = u.Update(c, db.GetDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	return models.FindUser(c, db.GetDB(), id)
}
