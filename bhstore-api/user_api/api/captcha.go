package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		zap.S().Errorw("生成验证码错误:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "生成验证码失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":      1,
		"captchaId": id,
		"picPath":   b64s,
	})
}
