package thelper

import (
	"context"
	"graphql/db"
	"graphql/middleware"
	"graphql/models"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/99designs/gqlgen/client"
)

func AddContext(t *testing.T) client.Option {
	t.Helper()
	u, err := insertAuthenticatedUser()
	if err != nil {
		return nil
	}

	ginCtx := setGinCtx(u)

	return func(bd *client.Request) {
		ginCtx.Request = bd.HTTP
		rawCtx := context.WithValue(ginCtx.Request.Context(), middleware.GinCtxKey, ginCtx)
		bd.HTTP = bd.HTTP.WithContext(rawCtx)
	}
}

func setGinCtx(u *models.User) *gin.Context {
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	ginContext.Set(middleware.GinCtxKey, u)

	return ginContext
}

func insertAuthenticatedUser() (*models.User, error) {
	// 認証に使う user を DB に入れる
	u := models.User{
		UserID:  "authenticatedUser99999999",
		Email:   "authenticatedUser%d@example.com",
		Name:    "authenticated user for test",
		Picture: "https://images.unsplash.com/photo-1485546246426-74dc88dec4d9?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1769&q=80",
	}

	err := u.Insert(context.Background(), db.GetDB(), boil.Infer())
	if err != nil {
		return nil, err
	}

	return &u, nil
}
