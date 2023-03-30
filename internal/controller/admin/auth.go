package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paper/config"
	"paper/internal/params"
	"paper/internal/service"
	"paper/pkg/sessions"
)

type Auth struct {
	AuthService service.AuthService
}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) GetLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"isNoLogin": true,
	})
}

func (a Auth) Longin(ctx *gin.Context) {
	var login params.LoginParam

	err := ctx.ShouldBind(&login)
	if err != nil {
		config.Logger.Error("请求参数格式错误：%s", err.Error())
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"isNoLogin": true,
			"message":   "请输入正确的账号密码！",
		})
		return
	}

	adminUser, err := a.AuthService.Login(login.Name, login.Password)
	if err != nil {
		config.Logger.Error("用户不存在：%s", err.Error())
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"isNoLogin": true,
			"message":   "请输入正确的账号密码！",
		})
		return
	}

	config.Logger.Info("adminUser：%s", adminUser)
	sessions.Save(ctx, adminUser.Uuid)
	ctx.Redirect(http.StatusMovedPermanently, "/admin")
}

func (a Auth) Logout(ctx *gin.Context) {
	var login params.LoginParam

	err := ctx.BindJSON(&login)
	if err != nil {

	}
}
