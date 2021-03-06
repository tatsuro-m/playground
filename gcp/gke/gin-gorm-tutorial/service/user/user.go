package user

import (
	"gin-gorm-tutorial/db"
	"gin-gorm-tutorial/entity"

	"github.com/gin-gonic/gin"
)

type Service struct{}

type User entity.User

func (s Service) GetAll() ([]User, error) {
	d := db.GetDB()
	var u []User

	if err := d.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (s Service) CreateModel(c *gin.Context) (User, error) {
	d := db.GetDB()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	if err := d.Create(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s Service) GetByID(id string) (User, error) {
	d := db.GetDB()
	var u User

	if err := d.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s Service) UpdateByID(id string, c *gin.Context) (User, error) {
	d := db.GetDB()
	var u User

	if err := d.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	d.Save(&u)

	return u, nil
}

func (s Service) DeletedByID(id string) error {
	d := db.GetDB()
	var u User

	if err := d.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
