package initialize

import (
	"bhstore/bhstore-api/user_api/global"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
)

func InitValidatorTrans(local string) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		var o bool
		global.Trans, o = uni.GetTranslator(local)
		if !o {
			zap.S().Errorf("uni.GetTranslator(%s) failed", local)
		}
		switch local {
		case "en":
			_ = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			_ = chTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			_ = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
	}
}
