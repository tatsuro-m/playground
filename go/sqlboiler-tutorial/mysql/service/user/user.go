package user

import (
	"context"
	"errors"
	"math/rand"
	"sqlboiler-tutorial-mysql/db"
	"sqlboiler-tutorial-mysql/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
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

func Insert() {
	random, err := MakeRandomStr(10)
	if err != nil {
		return
	}

	u := models.User{Name: "test", Email: random + "@example.com"}
	u.Insert(context.Background(), db.GetDB(), boil.Infer())
}

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

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
func GetPostsByUserID(userID int) ([]*models.Post, error) {
	u, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return u.Posts().All(context.Background(), db.GetDB())
}
