package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"log"
	"net/http"
	"os"
	"pcode/pkg/db"
	"pcode/pkg/graph/generated"
	"pcode/pkg/graph/resolver"
	"pcode/pkg/util"
	"time"
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

func loggingReq(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stdout, "[%s][%s]: %s\n", r.Method, r.URL, time.Now().Format(time.RFC3339))
		next.ServeHTTP(w, r)
	}
}

func main() {
	err := db.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	http.HandleFunc("/query", loggingReq(graphqlHandler))
	if !util.IsProd() {
		http.HandleFunc("/", loggingReq(playgroundHandler))
	}

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
