package api

import (
	"bhstore/bhstore-api/user_api/forms"
	"bhstore/bhstore-api/user_api/global"
	"bhstore/bhstore-api/user_api/middleware"
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

	_, err = global.UserClient.GetUserByMobile(c, &proto.MobileRequest{Mobile: registerInfo.Mobile})
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用户已存在",
		})
	}

	resp1, err := global.UserClient.CreateUser(c, &proto.CreateUserInfo{
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
		"msg":  resp1,
	})
}
func Login(c *gin.Context) {
	loginInfo := &forms.LoginForm{}
	//表单验证
	err := c.ShouldBind(loginInfo)
	if err != nil {
		utils.HandleValidator(err, c)
		return
	}
	//查找是否存在用户
	resp, err := global.UserClient.GetUserByMobile(c, &proto.MobileRequest{Mobile: loginInfo.Mobile})
	if err != nil {
		utils.HandleGrpcErrorToHttp(err, c)
		return
	}
	//验证密码
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
		return
	}
	//校对验证码
	ok := store.Verify(loginInfo.CaptchaId, loginInfo.Captcha, true)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "验证码错误",
		})
		return
	}
	//分发token
	token, err := middleware.GenToken(resp.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "生成token失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  token,
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
