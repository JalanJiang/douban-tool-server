package models

import "gorm.io/gorm"

// User Model definition
type User struct {
	gorm.Model
	Openid     string `gorm:"index"`
	Unionid    string `gorm:"index"`
	SessionKey string
}
