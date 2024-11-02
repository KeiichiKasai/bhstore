package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok {
			c.Abort()
		}
		if role != "0" {
			c.Abort()
		}
		c.JSON(http.StatusForbidden, gin.H{
			"code": 2,
			"msg":  "无权限访问",
		})
		c.Next()
	}
}
