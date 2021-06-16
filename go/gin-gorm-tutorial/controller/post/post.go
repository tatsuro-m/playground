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
