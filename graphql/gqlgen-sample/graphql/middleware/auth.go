package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getBearerToken(c.Request.Header.Get("Authorization"))

		// トークンが正しく設定されていないならそのまま Next する。
		if token == "" {
			fmt.Println("トークンが正しく設定されていません")
			c.Next()
			return
		}

		// https://github.com/99designs/gqlgen/blob/master/docs/content/recipes/authentication.md
		// トークンを検証してからこちらに書いてあるようなことは一通りやりたい。

		c.Next()
	}
}

func getBearerToken(header string) string {
	bearer := "Bearer "
	if header == "" || strings.Contains(header, bearer) == false {
		return ""
	}

	return strings.Replace(header, bearer, "", 1)
}
