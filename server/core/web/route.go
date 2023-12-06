package web

import (
	"github.com/gin-gonic/gin"
	controller2 "go_archetype/core/web/controller"
)

func initRoute(r *gin.Engine) {
	// ping
	r.GET("ping/v1", controller2.Ping.V1)
	r.GET("ping/v2", controller2.Ping.V2)

	// user
	r.GET("user/detail", controller2.UserCtrl.GetDetail)
	r.GET("user/list", controller2.UserCtrl.List)
	r.POST("user/create", controller2.UserCtrl.Create)
	r.POST("user/batch_create", controller2.UserCtrl.BatchCreate)
}
