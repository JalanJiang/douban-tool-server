package models

import "gorm.io/gorm"

// UserSubscription definition
type UserSubscription struct {
	gorm.Model
	UserID        uint `gorm:"index"`
	GroupID       string
	GroupName     string
	SearchKeyword string
}
