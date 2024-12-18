package main

import (
	"bhstore/bhstore-srv/user_srv/handler"
	"bhstore/bhstore-srv/user_srv/initialize"
	"bhstore/bhstore-srv/user_srv/proto"
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "server ip")
	Port := flag.Int("port", 8888, "server port")

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	flag.Parse()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserService{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		zap.S().Errorf("failed to listen: %v", err)
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
