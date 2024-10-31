package api

import (
	"bhstore/bhstore-api/user_api/forms"
	"bhstore/bhstore-api/user_api/global"
	"bhstore/bhstore-api/user_api/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	registerInfo := &forms.RegisterForm{}
	err := c.ShouldBindJSON(registerInfo)
	if err != nil {
		//TODO 验证器需要实现
	}
	resp, err := global.UserClient.CreateUser(c, &proto.CreateUserInfo{
		Nickname: registerInfo.Nickname,
		Mobile:   registerInfo.Mobile,
		Password: registerInfo.PassWord,
	})
	if err != nil {
		c.JSON(500, "失败")
	}
	c.JSON(200, resp)
}
func Login(c *gin.Context) {
	loginInfo := &forms.LoginForm{}
	err := c.ShouldBindJSON(loginInfo)
	if err != nil {
		//TODO 验证器
	}
	resp, err := global.UserClient.GetUserByMobile(c, &proto.MobileRequest{Mobile: loginInfo.Mobile})
	if err != nil {
		c.JSON(500, err)
	}

	resp2, err := global.UserClient.CheckPassWord(c, &proto.PasswordCheckInfo{
		Password:   loginInfo.PassWord,
		EnPassword: resp.Password,
	})
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, resp2)
}
func GetUserList(c *gin.Context) {
	resp1, err := global.UserClient.GetUserList(c, &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, resp1)
}
