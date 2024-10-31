package api

import (
	"bhstore/bhstore-api/user_api/forms"
	"bhstore/bhstore-api/user_api/global"
	"bhstore/bhstore-api/user_api/proto"
	"bhstore/bhstore-api/user_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	registerInfo := &forms.RegisterForm{}
	err := c.ShouldBind(registerInfo)
	if err != nil {
		utils.HandleValidator(err, c)
		return
	}
	resp, err := global.UserClient.CreateUser(c, &proto.CreateUserInfo{
		Nickname: registerInfo.Nickname,
		Mobile:   registerInfo.Mobile,
		Password: registerInfo.PassWord,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  resp,
	})
}
func Login(c *gin.Context) {
	loginInfo := &forms.LoginForm{}
	err := c.ShouldBind(loginInfo)
	if err != nil {
		utils.HandleValidator(err, c)
		return
	}
	resp, err := global.UserClient.GetUserByMobile(c, &proto.MobileRequest{Mobile: loginInfo.Mobile})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}
	resp2, err := global.UserClient.CheckPassWord(c, &proto.PasswordCheckInfo{
		Password:   loginInfo.PassWord,
		EnPassword: resp.Password,
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}
	if !resp2.Ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "密码错误",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  resp2,
	})
}
func GetUserList(c *gin.Context) {
	pn, _ := strconv.Atoi(c.DefaultQuery("pn", "1"))
	pSize, _ := strconv.Atoi(c.DefaultQuery("pSize", "5"))

	resp, err := global.UserClient.GetUserList(c, &proto.PageInfo{
		Pn:    int32(pn),
		PSize: int32(pSize),
	})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}
	ret := make([]interface{}, 0)
	for _, v := range resp.Data {
		data := make(map[string]interface{})

		data["id"] = v.Id
		data["name"] = v.Nickname
		data["mobile"] = v.Mobile
		data["role"] = v.Role

		ret = append(ret, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  ret,
	})
}
