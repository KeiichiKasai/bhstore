package main

import (
	"bhstore/bhstore-srv/user_srv/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func main() {
	dsn := "root:zx1913683154@tcp(127.0.0.1:3306)/bh_srv?charset=utf8mb4&parseTime=True&loc=Local"
	MyLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: MyLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&model.User{})
}
