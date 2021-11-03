package resolver

import (
	"graphql/graph/generated"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
)

func createGqlClient(t *testing.T) *client.Client {
	t.Helper()
	return client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}})))
}
