package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"pcode/pkg/code"
	"pcode/pkg/graph"
	"pcode/pkg/graph/gqlmodel"
	"pcode/pkg/service/address"
	"strconv"
)

func (r *queryResolver) Address(ctx context.Context, postalCode string) (*gqlmodel.Address, error) {
	a, err := address.Service{}.GetAddress(postalCode)
	if err != nil {
		return nil, graph.NewGqlError(err.Error(), code.RecordNotFoundErr)
	}

	return &gqlmodel.Address{ID: strconv.Itoa(a.PostalCode.ID), Name: a.Name}, nil
}
