package thelper

import (
	"context"
	"fmt"
	"graphql/db"
	"graphql/models"
	"strconv"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func InsertUser(t *testing.T, num int) []models.User {
	t.Helper()

	d := db.GetDB()
	ctx := context.Background()
	users := make([]models.User, 0)

	for i := 0; i < num; i++ {
		u := models.User{
			UserID:  fmt.Sprintf("test%d", i),
			Email:   fmt.Sprintf("test%d@example.com", i),
			Name:    "test" + strconv.Itoa(i),
			Picture: "https://images.ctfassets.net/xsbsj4s4spn6/63gTAXfljeNEDb65PW4uOc/3691198754392411b00a9f7a20df78a1/2021_lal_mr_yearbook_cover_v1_jd-1328.jpg?q=80",
		}

		err := u.Insert(ctx, d, boil.Infer())
		if err != nil {
			return nil
		}
		users = append(users, u)
	}

	return users
}

func InsertPost(t *testing.T, num int, userID int) []models.Post {
	t.Helper()

	d := db.GetDB()
	ctx := context.Background()
	user, err := models.FindUser(ctx, d, userID)
	if err != nil {
		return nil
	}

	posts := make([]models.Post, 0)

	for i := 0; i < num; i++ {
		p := models.Post{Title: "test" + strconv.Itoa(i)}
		err := user.AddPosts(ctx, d, true, &p)

		if err != nil {
			return nil
		}
		posts = append(posts, p)
	}

	return posts
}
