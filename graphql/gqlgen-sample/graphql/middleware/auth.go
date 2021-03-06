package middleware

import (
	"context"
	"graphql/models"
	"graphql/service/user"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"

	firebase "firebase.google.com/go/v4"

	"github.com/gin-gonic/gin"
)

const UserCtxKey = "user"

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getBearerToken(c.Request.Header.Get("Authorization"))

		if token == "" {
			c.Next()
			return
		}

		verifiedToken, err := verifyIdToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		u := getUser(verifiedToken)
		c.Set(UserCtxKey, &u)

		c.Next()
	}
}

func getUser(verifiedToken *auth.Token) models.User {
	s := user.Service{}
	b := s.ExistsByUID(verifiedToken.UID)

	var resUser models.User
	if b {
		data, err := s.GetUserByUID(verifiedToken.UID)
		if err != nil {
			return models.User{}
		}

		resUser = *data
	} else {
		u := models.User{UserID: verifiedToken.UID, Email: verifiedToken.Claims["email"].(string), Name: verifiedToken.Claims["name"].(string), Picture: verifiedToken.Claims["picture"].(string)}

		data, err := s.CreateUser(u)
		if err != nil {
			return models.User{}
		}

		resUser = data
	}

	return resUser
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

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *models.User {
	raw, ok := ctx.Value(UserCtxKey).(*models.User)

	// キャストできないような値が入っていたのなら空の構造体の pointer を返してしまう
	if ok == false {
		return &models.User{}
	}

	return raw
}
