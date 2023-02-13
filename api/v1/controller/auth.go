package controller

import (
	"github.com/gin-gonic/gin"
	"paper/api/v1/params"
)

type Auth struct {
}

func NewAuth() Auth {
	return Auth{}
}

func (a Auth) Signup(ctx *gin.Context) {
	var register params.RegisterParam

	err := ctx.BindJSON(&register)
	if err != nil {

	}
}

func (a Auth) Longin(ctx *gin.Context) {
	var login params.LoginParam

	err := ctx.BindJSON(&login)
	if err != nil {

	}
}

func (a Auth) Longout(ctx *gin.Context) {
	var login params.LoginParam

	err := ctx.BindJSON(&login)
	if err != nil {

	}
}
