package models

import "gorm.io/gorm"

// User Model definition
type User struct {
	gorm.Model
	OpenID     string `gorm:"index"`
	UnionID    string `gorm:"index"`
	SessionKey string
}
