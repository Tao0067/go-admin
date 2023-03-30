package admin

import (
	"github.com/gin-gonic/gin"
)

type Home struct {
}

func NewHome() Home {
	return Home{}
}

func (h Home) Index(ctx *gin.Context) {
	ctx.HTML(200, "home.html", gin.H{})
}
