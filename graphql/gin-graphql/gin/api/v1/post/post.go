package post

import (
	"fmt"
	"gin/service/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

// GET /posts
func (ctrl Controller) Index(c *gin.Context) {
	posts, err := post.Service{}.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, posts)
	}
}
