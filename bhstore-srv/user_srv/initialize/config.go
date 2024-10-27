package initialize

import (
	"bhstore/bhstore-srv/user_srv/global"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	path := "user_srv/initialize/config.yaml"
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("读取文件失败")
		panic(err)
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println("反序列化失败")
		panic(err)
	}
	fmt.Println("InitConfig success!")
}
