package user

import (
	"context"
	"database/sql"
	"math/rand"
	"sqlboiler-tutorial/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Insert(d *sql.DB) {
	u := models.User{Name: "test", Email: RandomString(5) + "@example.com"}
	u.Insert(context.Background(), d, boil.Infer())
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
