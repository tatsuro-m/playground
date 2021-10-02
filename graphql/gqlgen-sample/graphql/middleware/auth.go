package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authorization")
		fmt.Println(c.Request.Header.Get("Authorization"))

		fmt.Println("authorization")
		fmt.Println(c.Request.Header.Get("authorization"))

		// https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/authentication.md
		// トークンを検証してからこちらに書いてあるようなことは一通りやりたい。

		c.Next()
	}
}
