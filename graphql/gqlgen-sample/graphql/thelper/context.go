package thelper

import (
	"context"
	"graphql/middleware"
	"graphql/models"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/client"
)

func SetUserToContext(t *testing.T, authenticatedUser *models.User) client.Option {
	t.Helper()

	ginCtx := setGinCtx(authenticatedUser)
	return getGqlClientOption(ginCtx)
}

func SetEmptyUserToContext(t *testing.T) client.Option {
	t.Helper()
	ginCtx := setGinCtx(&models.User{})
	return getGqlClientOption(ginCtx)
}

func getGqlClientOption(ginCtx *gin.Context) client.Option {
	return func(bd *client.Request) {
		ginCtx.Request = bd.HTTP
		rawCtx := context.WithValue(ginCtx.Request.Context(), middleware.GinCtxKey, ginCtx)
		bd.HTTP = bd.HTTP.WithContext(rawCtx)
	}
}

func setGinCtx(u *models.User) *gin.Context {
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	ginContext.Set(middleware.UserCtxKey, u)

	return ginContext
}
