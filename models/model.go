package models

import (
	"JalanJiang/douban-tool-server/pkg/setting"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// db 数据库句柄
var Db *gorm.DB

func init() {
	var (
		err      error
		dbName   string
		user     string
		password string
		host     string
		port     string
	)

	// 获取配置
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	port = sec.Key("PORT").String()

	// 初始化 db
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		port,
		user,
		password,
		dbName,
	)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("DB Connect error: %v", err)
	}

	// 创建表迁移
	Db.Migrator().CreateTable(&UserSubscription{})
}

// func CloseDB() {
// 	defer db.Close()
// }
