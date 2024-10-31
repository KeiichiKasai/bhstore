package global

import (
	"bhstore/bhstore-api/user_api/config"
	"bhstore/bhstore-api/user_api/proto"
)

var (
	UserClient  proto.UserClient
	SeverConfig *config.APIConfig
)
