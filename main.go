package main

import (
	"JalanJiang/douban-tool-server/migrations"
	"JalanJiang/douban-tool-server/pkg/setting"
	"JalanJiang/douban-tool-server/routers"
	"fmt"
	"net/http"
)

func main() {
	// 初始化路由
	router := routers.InitRouter()

	// 服务器设置
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// 数据库文件迁移
	migrations.Migrate()

	s.ListenAndServe()
}
