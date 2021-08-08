package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

// GET /ping
func (ct Controller) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping!",
	})
}
