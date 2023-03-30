package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"paper/internal/controller/admin"
	sessionsMiddlerware "paper/pkg/sessions"
)

func Router() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions(sessionsMiddlerware.SessionsCookieName, store))
	//r.Use(sessionsMiddlerware.SessonMiddlerware())

	auth := admin.NewAuth()
	r.GET("/login", auth.GetLogin)
	r.POST("/login", auth.Longin)
	r.POST("/logout", auth.Logout)

	AdminUser := admin.NewAdminUser()
	r.GET("/test", AdminUser.Index)

	adminRouter := r.Group("/admin", sessionsMiddlerware.SessonMiddlerware())
	{
		home := admin.NewHome()
		adminRouter.GET("/", home.Index)

		//AdminUser := admin.NewAdminUser()
		adminRouter.GET("/user", AdminUser.Index)
		adminRouter.POST("/user/add", AdminUser.Add)
		//adminRouter.POST("/detail", auth.Signup)
		//adminRouter.POST("/edit", auth.Signup)
		//adminRouter.POST("/del", auth.Signup)
	}
	return r
}
