package app

import (
	"github.com/gin-gonic/gin"
)

type Console gin.HandlerFunc
type Events gin.HandlerFunc
type Exceptions gin.HandlerFunc
type Https gin.HandlerFunc
type Jobs gin.HandlerFunc
type Listeners gin.HandlerFunc
type Models func(ctx *gin.Context)
type Providers gin.HandlerFunc
