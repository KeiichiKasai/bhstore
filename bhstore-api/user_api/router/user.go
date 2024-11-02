package router

import (
	"bhstore/bhstore-api/user_api/api"
	"bhstore/bhstore-api/user_api/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	u := Router.Group("/user")
	{
		u.GET("/register", api.Register)
		u.GET("/login", api.Login)
		u.GET("/list", middleware.JWT(), middleware.Admin(), api.GetUserList)
	}

}
