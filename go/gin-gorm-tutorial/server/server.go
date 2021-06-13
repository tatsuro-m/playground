package server

import (
	"fmt"
	user "gin-gorm-tutorial/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := router()
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}

func router() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	u := v1.Group("/users")

	{
		ctrl := user.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}
