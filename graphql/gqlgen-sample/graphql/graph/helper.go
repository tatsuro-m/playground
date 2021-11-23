package graph

import (
	"graphql/graph/gqlmodel"
	"graphql/modelconv"
	"graphql/models"
	"graphql/service/user"
)

func SetUser(p *models.Post) *gqlmodel.Post {
	gqlPost := modelconv.ModelToGqlPost(p)
	u, _ := user.Service{}.GetUserByID(p.UserID)
	gqlPost.User = modelconv.ModelToGqlUser(u)

	return gqlPost
}
