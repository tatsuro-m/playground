package test_helper

import (
	"gin-gorm-tutorial/db"
	"gin-gorm-tutorial/entity"
	"strconv"
	"testing"
	"time"
)

func TimeFormat(t *testing.T, time time.Time) string {
	t.Helper()

	// この形式にフォーマットして、json の time 型レスポンスの中に含まれているかをチェックする
	l := "2006-01-02"
	return time.Format(l)
}

func InsertUser(t *testing.T, num int) []entity.User {
	t.Helper()

	d := db.GetDB()
	users := make([]entity.User, 0)

	for i := 0; i < num; i++ {
		u := entity.User{FirstName: "first_name" + strconv.Itoa(i), LastName: "last_name" + strconv.Itoa(i)}
		d.Create(&u)
		users = append(users, u)
	}
	return users
}

func InsertPost(t *testing.T, num int, user entity.User) []entity.Post {
	t.Helper()

	d := db.GetDB()
	posts := make([]entity.Post, 0)

	for i := 0; i < num; i++ {
		p := entity.Post{Title: "title" + strconv.Itoa(i), Content: "content" + strconv.Itoa(i), UserID: user.ID}
		d.Create(&p)
		posts = append(posts, p)
	}
	return posts
}
