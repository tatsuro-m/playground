package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"pcode/pkg/code"
	"pcode/pkg/graph"
	"pcode/pkg/graph/gqlmodel"
	"pcode/pkg/service/postalcode"
	"strconv"
)

func (r *queryResolver) PostalCode(ctx context.Context, address string) (*gqlmodel.PostalCode, error) {
	p, err := postalcode.Service{}.GetOne(address)
	if err != nil {
		return nil, graph.NewGqlError(err.Error(), code.RecordNotFoundErr)
	}

	return &gqlmodel.PostalCode{ID: strconv.Itoa(p.ID), Code: p.Code}, nil
}
