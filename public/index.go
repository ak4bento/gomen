// @Title
// @Description
// @Author
// @Update
package public

import (
	"gomen/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	app := gin.Default()

	routes.Run(app)

	app.Run()
}
