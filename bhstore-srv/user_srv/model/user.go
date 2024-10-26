package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int32          `gorm:"primary_key"`
	CreatedAt time.Time      `gorm:"column:add_time"`
	UpdatedAt time.Time      `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	IsDeleted bool
}

type User struct {
	BaseModel
	Mobile   string `gorm:"index:idx_mobile;unique;not null;type:varchar(11)"`
	Password string `gorm:"type:varchar(100);not null"`
	Nickname string `gorm:"type:varchar(20);not null"`
	Role     int32  `gorm:"default:1"`
}
