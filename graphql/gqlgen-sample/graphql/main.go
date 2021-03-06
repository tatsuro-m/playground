package main

import (
	"graphql/db"
	"graphql/graph"
	"graphql/graph/generated"
	"graphql/graph/resolver"
	"graphql/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	c := generated.Config{Resolvers: &resolver.Resolver{}}
	graph.ConfigDirectives(&c)
	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
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
	err := db.Init()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Setting up Gin
	r := gin.Default()

	r.Use(middleware.Cors())
	r.Use(middleware.Authentication())
	r.Use(middleware.GinContextToContext())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run(":8080")
}
