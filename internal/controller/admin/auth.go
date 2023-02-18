package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paper/config"
	"paper/internal/params"
	"paper/internal/service"
)

type Auth struct {
	AuthService service.AuthService
}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) Index(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{})
}

func (a Auth) Signup(ctx *gin.Context) {
	var register params.RegisterParam

	err := ctx.BindJSON(&register)
	if err != nil {
		config.Logger.Error("请求参数格式错误：%s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
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

	err := ctx.BindJSON(&login)
	if err != nil {

	}
}

func (a Auth) Logout(ctx *gin.Context) {
	var login params.LoginParam

	err := ctx.BindJSON(&login)
	if err != nil {

	}
}
