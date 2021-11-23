package resolver

import (
	"graphql/graph"
	"graphql/graph/generated"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
)

func createGqlClient(t *testing.T) *client.Client {
	t.Helper()

	c := generated.Config{Resolvers: &Resolver{}}
	graph.ConfigDirectives(&c)
	return client.New(handler.NewDefaultServer(generated.NewExecutableSchema(c)))
}
