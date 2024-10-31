package router

import (
	"bhstore/bhstore-api/user_api/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	u := Router.Group("/user")
	u.GET("/register", api.Register)
	u.GET("/login", api.Login)
	u.GET("/list", api.GetUserList)
}
