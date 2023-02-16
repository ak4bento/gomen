// @Title
// @Description
// @Author
// @Update
package routes

import (
	"github.com/gin-gonic/gin"
	"gomen/app/http"
)

func Run(app *gin.Engine) {
	var (
		controller http.Controller
	)

	app.GET("/index", controller.Index)
	app.GET("/index/show", controller.Show)
}
