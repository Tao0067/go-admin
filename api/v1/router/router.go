package router

import (
	"github.com/gin-gonic/gin"
	"paper/api/v1/controller"
)

func Router() *gin.Engine {
	r := gin.Default()
	auth := controller.NewAuth()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/admin/login", auth.Longin)
		v1.POST("/admin/signup", auth.Signup)
		v1.POST("/admin/logout", auth.Longout)
	}
	return r
}
