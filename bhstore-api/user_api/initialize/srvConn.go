package initialize

import (
	"bhstore/bhstore-api/user_api/global"
	"bhstore/bhstore-api/user_api/proto"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitConn() {
	info := global.SeverConfig.UserSrvInfo
	target := fmt.Sprintf("%s:%s", info.Host, info.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorf("err:%v", err)
		panic(err)
	}
	global.UserClient = proto.NewUserClient(conn)
	fmt.Println("InitConn success!")
}
