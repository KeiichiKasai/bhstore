package global

import (
	"bhstore/bhstore-srv/user_srv/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config config.Config
)
