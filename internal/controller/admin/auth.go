package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"paper/config"
	"paper/internal/params"
	"paper/internal/service"
)

type Auth struct {
	AuthService service.AuthService
}

const DefaultAuthInfoKye = "AUTH"

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) Index(ctx *gin.Context) {
	session := sessions.Default(ctx)
	admin := session.Get(DefaultAuthInfoKye)

	if admin == nil {
		ctx.Redirect(http.StatusMovedPermanently, "/admin/login")
		return
	}

	ctx.HTML(200, "home.html", gin.H{})
}

func (a Auth) GetLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	admin := session.Get(DefaultAuthInfoKye)

	if admin != nil {
		ctx.Redirect(http.StatusMovedPermanently, "/admin")
		return
	}

	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"isNoLogin": true,
	})
}

func (a Auth) Signup(ctx *gin.Context) {
	var register params.RegisterParam

	err := ctx.BindJSON(&register)
	if err != nil {
		config.Logger.Error("请求参数格式错误：%s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请求参数格式错误!",
		})
		return
	}

	err = a.AuthService.Signup(register.Name, register.Password)

	if err != nil {
		config.Logger.Error("账号密码错误：%s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "ok",
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

	session := sessions.Default(ctx)
	session.Set(DefaultAuthInfoKye, adminUser.Uuid)
	err = session.Save()
	ctx.Redirect(http.StatusMovedPermanently, "/admin")
}

func (a Auth) Logout(ctx *gin.Context) {
	var login params.LoginParam

	err := ctx.BindJSON(&login)
	if err != nil {

	}
}
