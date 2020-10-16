package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	// Cfg 配置读取
	Cfg *ini.File

	// RunMode 运行模式
	RunMode string

	// HTTPPort HTTP 端口
	HTTPPort int
	// ReadTimeout 读超时
	ReadTimeout time.Duration
	// WriteTimeout 写超时
	WriteTimeout time.Duration
)

func init() {
	var err error
	// 加载配置文件
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	loadServer()
}

func loadServer() {
	// 读取服务器配置
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(69)) * time.Second
}
