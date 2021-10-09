package modelconv

import (
	"graphql/graph/gqlmodel"
	"graphql/models"
	"strconv"
)

func ModelToGqlPost(mp *models.Post) *gqlmodel.Post {
	return &gqlmodel.Post{ID: strconv.Itoa(mp.ID), Title: mp.Title, CreatedAt: mp.CreatedAt, UpdatedAt: mp.UpdatedAt}
}

func ModelToGqlUser(mu *models.User) *gqlmodel.User {
	return &gqlmodel.User{ID: strconv.Itoa(mu.ID), Email: mu.Email, Name: mu.Name, Picture: mu.Picture, CreatedAt: mu.CreatedAt, UpdatedAt: mu.UpdatedAt}
}
