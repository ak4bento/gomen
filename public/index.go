// @Title
// @Description
// @Author
// @Update
package public

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gomen/app/providers/orm"
	"gomen/routes"
)

var (
	DB orm.Model
)

func init() {
	app := gin.Default()

	routes.Run(app)

	app.Run(fmt.Sprintf("%s:%d", "0.0.0.0", 1712))
}
