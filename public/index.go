// @Title
// @Description
// @Author
// @Update
package public

import (
	"github.com/gin-gonic/gin"
	"gomen/routes"
)

func init() {
	app := gin.Default()

	routes.Run(app)

	app.Run()
}
