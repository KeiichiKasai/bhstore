package handler

import (
	"bhstore/bhstore-srv/user_srv/global"
	"bhstore/bhstore-srv/user_srv/model"
	"bhstore/bhstore-srv/user_srv/proto"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserService struct{}

func Paginate(page, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 10:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserService) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	var userList []model.User
	result := global.DB.Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	response := proto.UserListResponse{}
	result = global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&userList)
	for _, user := range userList {
		temp := proto.UserInfoResponse{
			Id:       user.ID,
			Password: user.Password,
			Mobile:   user.Mobile,
			Nickname: user.Nickname,
			Role:     user.Role,
		}
		response.Data = append(response.Data, &temp)
		response.Total = int32(result.RowsAffected)
	}
	return &response, nil
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
	user = model.User{
		Mobile:   req.Mobile,
		Nickname: req.Nickname,
		Password: req.Password,
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
	
}
