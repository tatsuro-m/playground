package main

import (
	"gin/logging"
	"os"
	"time"

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
	})

	sugar := logging.GetS()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Error("エラー")

	r.Run(":" + os.Getenv("PORT"))
}
