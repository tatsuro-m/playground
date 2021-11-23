package main

import (
	"gin/logging"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	l := logging.InitLogger()
	defer l.Sync()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "cloud run!!",
		})

		logging.S().Infow("success!!")
	})

	r.GET("/error", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)

		logging.S().Error("test not found error",
			"status", http.StatusNotFound,
		)
	})

	r.Run(":" + os.Getenv("PORT"))
}
