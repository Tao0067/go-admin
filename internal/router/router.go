package router

import (
	"github.com/gin-gonic/gin"
	"paper/internal/controller/admin"
)

func Router() *gin.Engine {
	r := gin.Default()
	auth := admin.NewAuth()

	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/index", auth.Index)
		adminRouter.POST("/login", auth.Longin)
		adminRouter.POST("/signup", auth.Signup)
		adminRouter.POST("/logout", auth.Logout)
	}
	return r
}
