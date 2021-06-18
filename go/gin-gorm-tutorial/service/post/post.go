package post

import (
	"gin-gorm-tutorial/db"
	"gin-gorm-tutorial/entity"
)

type Service struct{}
type Post entity.Post

func (s Service) GetAll() ([]Post, error) {
	d := db.GetDB()
	var p []Post

	if err := d.Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}
