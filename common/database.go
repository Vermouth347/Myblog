package common

import (
	"fmt"
	"myblog/model"
	"net/url"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// 数据库初始化
func InitDB() *gorm.DB {
	driverName := "mysql"
	user := "root"
	password := "123456"
	host := "localhost"
	port := "3306"
	database := "blog"
	charset := "utf8"
	loc := "Asian/Fukuoka"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		user,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	// 连接数据库
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	// 迁移数据表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

// 数据库信息获取
func GetDB() *gorm.DB {
	return DB
}
