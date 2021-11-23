package graph

import (
	"context"
	"errors"
	"graphql/code"
	"graphql/ginctx"
	"graphql/graph/generated"
	"graphql/graph/gqlmodel"
	"graphql/modelconv"
	"graphql/models"
	"graphql/service/user"

	"github.com/99designs/gqlgen/graphql"
)

func SetUser(p *models.Post) *gqlmodel.Post {
	gqlPost := modelconv.ModelToGqlPost(p)
	u, _ := user.Service{}.GetUserByID(p.UserID)
	gqlPost.User = modelconv.ModelToGqlUser(u)

	return gqlPost
}

func ConfigDirectives(c *generated.Config) {
	c.Directives.Authenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		_, err := ginctx.GetUserFromGinCtx(ctx)
		if err != nil {
			return nil, errors.New(code.NotAuthorize)
		}

		return next(ctx)
	}
}
