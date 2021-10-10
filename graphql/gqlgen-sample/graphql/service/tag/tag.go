package tag

import (
	"context"
	"graphql/db"
	"graphql/models"
)

type Service struct{}

func (s Service) GetByID(id int) (*models.Tag, error) {
	return models.FindTag(context.Background(), db.GetDB(), id)
}
