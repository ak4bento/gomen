package models

import (
	"github.com/gin-gonic/gin"
)

type Home struct {
	Id       string `form:"id"`
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}

func (this *Home) TableName() string {
	return "app_home"
}

func (m *Home) Request(ctx *gin.Context) {
	ctx.ShouldBind(&m)
	return
}
