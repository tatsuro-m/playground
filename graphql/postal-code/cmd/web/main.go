package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"pcode/db"
	"pcode/graph/generated"
	"pcode/graph/resolver"
	"pcode/util"
)

func graphqlHandler() gin.HandlerFunc {
	c := generated.Config{Resolvers: &resolver.Resolver{}}
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
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Setting up Gin
	r := gin.Default()

	r.POST("/query", graphqlHandler())
	if !util.IsProd() {
		r.GET("/", playgroundHandler())
	}
	r.Run(":8080")
}
