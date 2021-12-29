package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	generated1 "pcode/pkg/graph/generated"
	gqlmodel1 "pcode/pkg/graph/gqlmodel"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input gqlmodel1.NewTodo) (*gqlmodel1.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*gqlmodel1.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
