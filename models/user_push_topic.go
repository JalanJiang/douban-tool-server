package models

import "gorm.io/gorm"

// UserPushTopic definition
type UserPushTopic struct {
	gorm.Model
	UserID      uint `gorm:"unique"`
	LastTopicID uint
}
