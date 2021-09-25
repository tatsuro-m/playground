package post

import (
	"fmt"
	"gin/db"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

// GET /posts
func (ctrl Controller) Index(c *gin.Context) {
	posts, err := models.Posts().All(c, db.GetDB())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, posts)
	}
}
