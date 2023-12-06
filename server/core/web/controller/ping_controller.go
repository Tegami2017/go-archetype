package controller

import "github.com/gin-gonic/gin"

type PingController struct {
	name *string
}

var Ping = &PingController{}

func (con PingController) V1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (con PingController) V2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong2",
	})
}
