package middleware

import (
	"context"
	"graphql/models"

	"github.com/gin-gonic/gin"
)

func GinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *models.User {
	// panic になるかもしれないけど、gin が Recover して 500 を返してくれるのでOK。
	raw := ctx.Value(userCtxKey).(*models.User)

	return raw
}
