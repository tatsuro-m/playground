package post

import (
	"fmt"
	"gin-gorm-tutorial/service/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

// GET /posts
func (ct Controller) Index(c *gin.Context) {
	var s post.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// POST /posts
func (ct Controller) Create(c *gin.Context) {
	var s post.Service

	if p, err := s.CreateModel(c); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusCreated, p)
	}
}

// GET /posts/:id
func (ct Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s post.Service

	p, err := s.GetByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// PUT /posts/:id
func (ct Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s post.Service

	p, err := s.UpdateByID(id, c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// DELETE /posts/:id
func (ct Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s post.Service

	p, err := s.DeleteByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}
