package initialize

import (
	"fmt"
	"go.uber.org/zap"
)

func InitLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("InitLogger err:", err)
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	fmt.Println("InitLogger success!")
}
