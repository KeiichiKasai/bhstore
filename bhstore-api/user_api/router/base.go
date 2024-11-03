package router

import (
	"bhstore/bhstore-api/user_api/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	b := Router.Group("/base")
	{
		b.GET("/captcha", api.Captcha)
	}
}
