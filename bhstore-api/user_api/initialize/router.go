package initialize

import (
	"bhstore/bhstore-api/user_api/middleware"
	"bhstore/bhstore-api/user_api/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	fmt.Println("InitRouters success!")

	r := gin.Default()
	r.Use(middleware.CORS())
	APIRouter := r.Group("/bh_store")
	router.InitUserRouter(APIRouter)

	return r
}
