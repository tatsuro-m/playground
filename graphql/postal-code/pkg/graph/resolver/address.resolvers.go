package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"pcode/pkg/graph/gqlmodel"
)

func (r *queryResolver) Address(ctx context.Context, postalCode int) (*gqlmodel.Address, error) {
	panic(fmt.Errorf("not implemented"))
}
