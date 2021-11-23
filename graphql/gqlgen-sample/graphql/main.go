package main

import (
	"context"
	"errors"
	"graphql/code"
	"graphql/db"
	"graphql/ginctx"
	"graphql/graph/generated"
	"graphql/graph/resolver"
	"graphql/middleware"

	"github.com/99designs/gqlgen/graphql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func graphqlHandler() gin.HandlerFunc {
	c := generated.Config{Resolvers: &resolver.Resolver{}}
	configDirectives(&c)
	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func configDirectives(c *generated.Config) {
	c.Directives.Authenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		_, err := ginctx.GetUserFromGinCtx(ctx)
		if err != nil {
			return nil, errors.New(code.NotAuthorize)
		}

		return next(ctx)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	db.Init()
	defer db.Close()

	// Setting up Gin
	r := gin.Default()

	r.Use(middleware.Cors())
	r.Use(middleware.Authentication())
	r.Use(middleware.GinContextToContext())

	r.GET("/ping", func(c *gin.Context) {
		u := middleware.ForContext(c)

		c.JSON(200, gin.H{
			"user": u,
		})
	})

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":8080")
}
