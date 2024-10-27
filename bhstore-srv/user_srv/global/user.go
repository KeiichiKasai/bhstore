package global

import (
	"bhstore/bhstore-srv/user_srv/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config config.Config
	Log    zap.Logger
)
