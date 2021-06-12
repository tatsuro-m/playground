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
	u := r.Group("/users")

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
