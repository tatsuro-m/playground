package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

const GinCtxKey = "GinContextKey"

func GinContextToContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinCtxKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
