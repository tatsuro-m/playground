package tag

import (
	"context"
	"fmt"
	"graphql/db"
	"graphql/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Service struct{}

func (s Service) GetByID(id int) (*models.Tag, error) {
	return models.FindTag(context.Background(), db.GetDB(), id)
}

func (s Service) Posts(tagID int) (models.PostSlice, error) {
	posts, err := models.Posts(
		qm.InnerJoin("post_tags pt on id = pt.post_id"),
		qm.InnerJoin("tags t on pt.tag_id = t.id"),
		qm.Where("t.id = ?", tagID),
	).All(context.Background(), db.GetDB())

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return posts, nil
}
