package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"

	firebase "firebase.google.com/go/v4"

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

		verifiedToken, err := verifyIdToken(token)
		fmt.Println(verifiedToken)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

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

func verifyIdToken(token string) (*auth.Token, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	return client.VerifyIDToken(ctx, token)
}
