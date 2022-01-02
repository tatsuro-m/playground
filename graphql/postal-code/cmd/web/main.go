package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"pcode/pkg/db"
	"pcode/pkg/graph/generated"
	"pcode/pkg/graph/resolver"
	"pcode/pkg/util"
)

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c := generated.Config{Resolvers: &resolver.Resolver{}}
		h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

		h.ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Not support http method %s", r.Method)))
	}
}

func playgroundHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h := playground.Handler("GraphQL", "/query")
		h.ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Not support http method %s", r.Method)))
	}
}

func main() {
	err := db.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	http.HandleFunc("/query", graphqlHandler)
	if !util.IsProd() {
		http.HandleFunc("/", playgroundHandler)
	}

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
