package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paper/config"
	"paper/internal/model"
	"paper/internal/params"
	"paper/internal/service"
)

type AdminUser struct {
	AuthService service.AuthService
	AdminUser   service.AdminUserService
}

var users = []model.Admin{
	{Id: 1, Name: "Alice", Uuid: "20", CreateTime: 2331111},
	{Id: 2, Name: "Bob", Uuid: "30", CreateTime: 3312321332},
}

func NewAdminUser() AdminUser {
	return AdminUser{}
}

func (a AdminUser) Index(ctx *gin.Context) {
	var search params.SearchAdminUser
	err := ctx.ShouldBindQuery(&search)
	if err != nil {
		config.Logger.Error("请求参数格式错误：%s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请求参数格式错误!",
		})
	}
	if search.Page == 0 {
		search.Page = 1
	}

	if search.Size == 0 {
		search.Size = 5
	}
	list, total, err := a.AdminUser.AdminUserListPage(0, 0, search.Name, search.Page, search.Size)

	ctx.HTML(http.StatusOK, "user.html", gin.H{
		"users": list,
		"total": total,
	})
}

func (a AdminUser) Add(ctx *gin.Context) {
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

func (a AdminUser) Del(ctx *gin.Context) {
	var del params.DelAdminUser
	err := ctx.BindJSON(del)
	if err != nil {
		config.Logger.Error("请求参数格式错误：%s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请求参数格式错误!",
		})
		return
	}
	err = a.AdminUser.DelAdminUser(del.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请求参数格式错误!",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (a AdminUser) Edit(ctx *gin.Context) {
	var edit params.EditAdminUser
	err := ctx.BindJSON(&edit)
	if err != nil {
		config.Logger.Error("请求参数格式错误：%s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请求参数格式错误!",
		})
		return
	}
	err = a.AdminUser.EditAdminUser(edit.Id, edit.Name, edit.Password)
	if err != nil {
		config.Logger.Error("请求参数格式错误：%s", err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "请求参数格式错误!",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
