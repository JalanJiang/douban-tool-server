package migrations

import "JalanJiang/douban-tool-server/models"

// Migrate 执行迁移
func Migrate() {
	models.Db.Migrator().CreateTable(&models.UserSubscription{})
	models.Db.Migrator().CreateTable(&models.User{})
	models.Db.Migrator().CreateTable(&models.UserPushTopic{})
}
