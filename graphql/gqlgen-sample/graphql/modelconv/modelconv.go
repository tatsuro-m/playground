package modelconv

import (
	"graphql/graph/gqlmodel"
	"graphql/models"
	"strconv"
)

func ModelToGqlUser(mu *models.User) *gqlmodel.User {
	return &gqlmodel.User{ID: strconv.Itoa(mu.ID), Email: mu.Email, Name: mu.Name, Picture: mu.Picture, CreatedAt: mu.CreatedAt, UpdatedAt: mu.UpdatedAt}
}
