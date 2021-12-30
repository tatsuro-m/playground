package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"pcode/pkg/graph/gqlmodel"
)

func (r *queryResolver) Address(ctx context.Context, postalCode int) (*gqlmodel.Address, error) {
	return &gqlmodel.Address{}, nil
}
