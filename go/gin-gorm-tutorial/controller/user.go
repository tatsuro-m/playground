package user

import (
	"fmt"
	user "gin-gorm-tutorial/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

// GET /users
func (ct Controller) Index(c *gin.Context) {
	var s user.Service
	p, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// POST /users
func (ct Controller) Create(c *gin.Context) {
	var s user.Service
	p, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusCreated, p)
	}
}

// GET /users/:id
func (ct Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	p, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// PUT /users/:id
func (ct Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	p, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// DELETE /users/:id
func (ct Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service

	if err := s.DeletedByID(id); err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusNoContent, gin.H{"id #" + id: "deleted"})
	}
}
