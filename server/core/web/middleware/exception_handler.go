package middleware

import (
	"github.com/gin-gonic/gin"
)

func ExceptionHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(200, gin.H{
				"code": 1001,
				"msg":  err,
				"data": nil,
			})
		}
	}()

	c.Next()
}
