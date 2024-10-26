package main

import (
	"bhstore/bhstore-srv/user_srv/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	userClient proto.UserClient
	conn       *grpc.ClientConn
)

func Init() {
	var err error
	conn, err = grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		resp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			Nickname: fmt.Sprintf("user%d", i),
			Mobile:   fmt.Sprintf("1852312981%d", i),
			Password: "admin123",
		})
		if err != nil {
			fmt.Println(err)
			panic(err)
			return
		}
		fmt.Println(resp)
	}
}
func TestGetUserList() {
	resp1, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range resp1.Data {
		fmt.Println(user.Mobile, user.Nickname, user.Password)
		resp2, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:   "admin123",
			EnPassword: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp2.Ok)
	}
}
func TestGetUserByMobile() {
	resp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: "18523129812"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
func TestGetUserById() {
	resp, err := userClient.GetUserById(context.Background(), &proto.IdRequest{Id: 2})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
func main() {
	Init()
	//TestCreateUser()
	//TestGetUserList()
	//TestGetUserByMobile()
	//TestGetUserById()
	err := conn.Close()
	if err != nil {
		panic(err)
	}
}
