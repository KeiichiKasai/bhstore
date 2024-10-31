package initialize

import (
	"bhstore/bhstore-api/user_api/router"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	APIRouter := r.Group("/bh_store")
	router.InitUserRouter(APIRouter)
	return r
}
