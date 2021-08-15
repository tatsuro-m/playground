package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tatsuro-m/hackernews/internal/auth"

	"github.com/tatsuro-m/hackernews/graph"
	"github.com/tatsuro-m/hackernews/graph/generated"

	"github.com/go-chi/chi"
	database "github.com/tatsuro-m/hackernews/internal/pkg/db/migrations/mysql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(auth.MiddleWare())

	database.InitDB()
	server := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
