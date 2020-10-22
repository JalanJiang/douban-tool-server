package models

import "gorm.io/gorm"

// UserSubscription definition
type UserSubscription struct {
	UserID        uint `gorm:"index"`
	GroupID       string
	GroupName     string
	SearchKeyword string
	gorm.Model
}

// AddUserSubscription 新增用户订阅
func AddUserSubscription(userID uint, groupID string, groupName string, searchKeyword string) (*gorm.DB, UserSubscription) {
	userSubscription := UserSubscription{
		UserID:        userID,
		GroupID:       groupID,
		GroupName:     groupName,
		SearchKeyword: searchKeyword,
	}
	result := Db.Create(&userSubscription)

	return result, userSubscription
}

// ExistUserSubscriptionByID 查询订阅是否存在
func ExistUserSubscriptionByID(usID uint, userID uint) bool {
	var userSubscription UserSubscription
	Db.Where("id = ? AND user_id = ?", usID, userID).Find(&userSubscription)
	if userSubscription.ID > 0 {
		return true
	}

	return false
}

// GetUserSubscriptionByID 通过订阅 ID 获取用户订阅
func GetUserSubscriptionByID(usID uint, userID uint) (*gorm.DB, UserSubscription) {
	var userSubscription UserSubscription
	result := Db.Where("id = ? AND user_id = ?", usID, userID).First(&userSubscription)
	return result, userSubscription
}

// EditUserSubscription 更新订阅
func EditUserSubscription(usID uint, searchKeyword string) bool {
	Db.Model(&UserSubscription{}).Where("id = ?", usID).Update("search_keyword", searchKeyword)
	// TODO: 返回更新后的数据
	return true
}

// DeleteUserSubscriptionByID 通过订阅 ID 删除订阅
func DeleteUserSubscriptionByID(useID uint) {
	Db.Delete(&UserSubscription{}, useID)
}
