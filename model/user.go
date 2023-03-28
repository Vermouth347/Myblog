package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model         // 可自动添加id等基本字段
	UserName    string `gorm:"varchar(20);not null"`
	PhoneNumber string `gorm:"varchar(20);not null;unique"`
	Password    string `gorm:"size:255;not null"`
	Avatar      string `gorm:"size:255;not null"`
	Collects    Array  `gorm:"type:longtext"`
	Following   Array  `gorm:"type:longtext"`
	Fans        int    `gorm:"AUTO_INCREMENT"`
}

// 部分用户信息，便于将数据库的查询结果绑定到结构体上
type UserInfo struct {
	ID       uint   `json:"id"`
	Avatar   string `json:"avatar"`
	UserName string `json:"userName`
}
