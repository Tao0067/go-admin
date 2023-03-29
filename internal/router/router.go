package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"paper/internal/controller/admin"
)

func Router() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	auth := admin.NewAuth()

	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/", auth.Index)
		adminRouter.GET("/login", auth.GetLogin)
		adminRouter.POST("/login", auth.Longin)
		adminRouter.POST("/signup", auth.Signup)
		adminRouter.POST("/logout", auth.Logout)
	}
	return r
}
