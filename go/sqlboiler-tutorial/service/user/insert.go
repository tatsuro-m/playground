package user

import (
	"context"
	"errors"
	"math/rand"
	"sqlboiler-tutorial/db"
	"sqlboiler-tutorial/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

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
