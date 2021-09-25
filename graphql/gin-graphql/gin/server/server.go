package server

import (
	"fmt"
	"gin/api/v1/post"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := Router()
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

func Router() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")

	p := v1.Group("/posts")

	{
		ctrl := post.Controller{}
		p.GET("", ctrl.Index)
	}

	return r
}
