package ginctx

import (
	"context"
	"errors"
	"fmt"
	"graphql/middleware"
	"graphql/models"

	"github.com/gin-gonic/gin"
)

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}

	return gc, nil
}

func GetUserFromGinCtx(ctx context.Context) (*models.User, error) {
	ctx, _ = GinContextFromContext(ctx)
	user := middleware.ForContext(ctx)
	if user.ID == 0 {
		return nil, errors.New("user is not valid")
	}

	return user, nil
}
