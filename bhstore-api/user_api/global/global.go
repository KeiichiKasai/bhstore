package global

import (
	"bhstore/bhstore-api/user_api/config"
	"bhstore/bhstore-api/user_api/proto"
	ut "github.com/go-playground/universal-translator"
)

var (
	UserClient  proto.UserClient
	SeverConfig *config.APIConfig
	Trans       ut.Translator
)
