package post

import (
	"fmt"
	"gin-gorm-tutorial/db"
	"gin-gorm-tutorial/entity"

	"github.com/gin-gonic/gin"
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

func (s Service) GetAllByUserID(id string) ([]Post, error) {
	d := db.GetDB()
	fmt.Println(id)
	fmt.Println("user の id からその user が持っているタスクを全て返す関数を作る")

	var u entity.User
	var p []Post

	if err := d.Model(&u).Association("posts").Find(&p); err != nil {
		return p, nil
	}

	return p, nil
}

func (s Service) GetByID(id string) (Post, error) {
	d := db.GetDB()
	var p Post

	if err := d.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (s Service) CreateModel(c *gin.Context) (Post, error) {
	d := db.GetDB()
	var p Post

	if err := c.BindJSON(&p); err != nil {
		return p, err
	}

	if err := d.Create(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (s Service) UpdateByID(id string, c *gin.Context) (Post, error) {
	d := db.GetDB()
	var p Post

	if err := d.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}

	if err := c.BindJSON(&p); err != nil {
		return p, err
	}

	d.Save(&p)
	return p, nil
}

func (s Service) DeleteByID(id string) (Post, error) {
	d := db.GetDB()
	var p Post

	data := d.Where("id = ?", id).First(&p)
	if err := data.Delete(&p).Error; err != nil {
		return p, err
	}

	return p, nil
}
