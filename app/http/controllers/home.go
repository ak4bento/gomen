// @Title
// @Description
// @Author
// @Update
package controllers

import (
	"github.com/gin-gonic/gin"
	"gomen/app/models"
)

type Home struct {
	id  string
	req models.Home
}

func (home *Home) Index(ctx *gin.Context) {
	home.req.Request(ctx)
	ctx.JSON(200, home.req)
	return
}

func (home *Home) Show(ctx *gin.Context) {
}

func (home *Home) Update(ctx *gin.Context) {
}

func (home *Home) Delete(ctx *gin.Context) {
}
