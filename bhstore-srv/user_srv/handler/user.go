package handler

import (
	"bhstore/bhstore-srv/user_srv/global"
	"bhstore/bhstore-srv/user_srv/model"
	"bhstore/bhstore-srv/user_srv/proto"
	"context"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/gorm"
	"strings"
)

type UserService struct{}

func ModelToResp(user model.User) *proto.UserInfoResponse {
	res := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Mobile:   user.Mobile,
		Nickname: user.Nickname,
		Role:     user.Role,
	}
	return &res

}

func Paginate(page, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserService) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	resp := &proto.UserListResponse{}
	resp.Total = int32(result.RowsAffected)
	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
	for _, user := range users {
		userResp := ModelToResp(user)
		resp.Data = append(resp.Data, userResp)
	}
	return resp, nil
}
func (s *UserService) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	response := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Mobile:   user.Mobile,
		Nickname: user.Nickname,
		Role:     user.Role,
	}
	return &response, nil
}
func (s *UserService) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	response := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Mobile:   user.Mobile,
		Nickname: user.Nickname,
		Role:     user.Role,
	}
	return &response, nil
}
func (s *UserService) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where("mobile = ?", req.Mobile).First(&user)
	if result.RowsAffected != 0 {
		return nil, errors.New("already existed")
	}
	option := password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, enPasswd := password.Encode(req.Password, &option)
	passwd := fmt.Sprintf("%s$%s", salt, enPasswd)
	user = model.User{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Password: passwd,
	}
	global.DB.Create(&user)
	response := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Mobile:   user.Mobile,
		Nickname: user.Nickname,
		Role:     user.Role,
	}
	return &response, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	user = model.User{
		Nickname: req.Nickname,
	}
	global.DB.Save(&user)
	response := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		Mobile:   user.Mobile,
		Nickname: user.Nickname,
		Role:     user.Role,
	}
	return &response, nil
}
func (s *UserService) CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	option := password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	enPasswdInfo := strings.Split(req.EnPassword, "$")
	check := password.Verify(req.Password, enPasswdInfo[0], enPasswdInfo[1], &option)
	return &proto.CheckResponse{Ok: check}, nil
}
