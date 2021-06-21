package server

import (
	"fmt"
	"gin-gorm-tutorial/controller/api/v1/post"
	"gin-gorm-tutorial/controller/api/v1/user"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

const (
	id = "/:id"
)

func router() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	u := v1.Group("/users")

	{
		ctrl := user.Controller{}
		u.GET("", ctrl.Index)
		u.GET(id, ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT(id, ctrl.Update)
		u.DELETE(id, ctrl.Delete)
		u.GET(id+"/posts", ctrl.Posts)
	}

	p := v1.Group("/posts")

	{
		ctrl := post.Controller{}
		p.GET("", ctrl.Index)
		p.GET(id, ctrl.Show)
		p.POST("", ctrl.Create)
		p.PUT(id, ctrl.Update)
		p.DELETE(id, ctrl.Delete)
	}

	return r
}
