package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()

	// テスト用なので CORS は全許可
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + os.Getenv("PORT"))
}
