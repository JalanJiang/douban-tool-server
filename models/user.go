package models

import (
	"gorm.io/gorm"
)

// User Model definition
type User struct {
	gorm.Model
	OpenID     string `gorm:"index"`
	UnionID    string
	SessionKey string
}

// GetUserByOpenID 通过 open_id 获取用户信息
func GetUserByOpenID(openID string) (*gorm.DB, User) {
	var user User
	result := Db.Where("open_id = ?", openID).First(&user)
	return result, user
}

// AddUser 添加用户
func AddUser(openID string, unionID string, sessionKey string) (*gorm.DB, User) {
	user := User{
		OpenID:     openID,
		UnionID:    unionID,
		SessionKey: sessionKey,
	}
	result := Db.Create(&user)

	return result, user
}
