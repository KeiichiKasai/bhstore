package main

import (
	"bhstore/bhstore-api/user_api/global"
	"bhstore/bhstore-api/user_api/initialize"
	"fmt"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitConn()
	e := initialize.InitRouters()
	e.Run(fmt.Sprintf(":%s", global.SeverConfig.Port))
}
