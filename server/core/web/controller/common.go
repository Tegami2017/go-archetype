package controller

import (
	"github.com/gin-gonic/gin"
)

func success(c *gin.Context, res any) {
	body := gin.H{
		"code": 0,
		"msg":  "",
		"data": res,
	}
	c.JSON(200, body)
}
