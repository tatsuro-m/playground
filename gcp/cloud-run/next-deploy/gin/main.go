package main

import (
	"gin/logging"
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

		logging.S().Infow("success!!",
			"path", "/ping",
		)
	})

	r.Run(":" + os.Getenv("PORT"))
}
