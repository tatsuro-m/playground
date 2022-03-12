package main

import (
	"graphql/db"
	"graphql/graph/generated"
	"graphql/graph/resolver"
	"graphql/middleware"
	"graphql/util"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c := generated.Config{Resolvers: &resolver.Resolver{}}
		h := handler.NewDefaultServer(generated.NewExecutableSchema(c))
		h.ServeHTTP(w, r)
	default:
		util.NotSupportHTTPMethod(w, r)
	}
}

func playgroundHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h := playground.Handler("GraphQL", "/query")
		h.ServeHTTP(w, r)
	default:
		util.NotSupportHTTPMethod(w, r)
	}
}

func main() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/query", middleware.LoggingReq(graphqlHandler))
	http.HandleFunc("/", middleware.LoggingReq(playgroundHandler))
	log.Fatalln(http.ListenAndServe(":8080", nil))

	//r.Use(middleware.Cors())
	//r.Use(middleware.Authentication())
	//r.Use(middleware.GinContextToContext())

}
