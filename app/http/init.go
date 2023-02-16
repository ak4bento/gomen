package http

import (
	"fmt"
	"gomen/app/http/controllers"
)

type Controller struct {
	controllers.Home
}

func init() {
	fmt.Println("This is http init")
}
