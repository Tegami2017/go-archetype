package web

import (
	"github.com/gin-gonic/gin"
	"go_archetype/core/web/middleware"
)

var r = gin.Default()

func init() {
	r.Use(middleware.ExceptionHandler)
	initRoute(r)
}

func Listen() {
	r.Run(":8080")
}
