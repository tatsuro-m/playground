package graph

import (
	"context"
	"errors"
	"graphql/code"
	"graphql/ginctx"
	"graphql/graph/generated"

	"github.com/99designs/gqlgen/graphql"
)

func ConfigDirectives(c *generated.Config) {
	c.Directives.Authenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		_, err := ginctx.GetUserFromGinCtx(ctx)
		if err != nil {
			return nil, errors.New(code.NotAuthorize)
		}

		return next(ctx)
	}
}
