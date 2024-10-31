package initialize

import (
	"bhstore/bhstore-api/user_api/global"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	path := "user_api/initialize/config.yaml"
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		zap.S().Infof("InitConfig Read In Config Failed, %s", err.Error())
		panic(err)
	}

	if err := v.Unmarshal(&global.SeverConfig); err != nil {
		zap.S().Infof("InitConfig Unmarshal Failed, %s", err.Error())
		panic(err)
	}
	fmt.Println("InitConfig success!")
}
