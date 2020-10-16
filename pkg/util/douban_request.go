package util

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	// Cfg 配置读取
	Cfg *ini.File
	// Host 域名
	Host string
	// Key 加密使用 key 值
	Key string
	// APIKey apikey
	APIKey string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	sec, err := Cfg.GetSection("douban")
	if err != nil {
		log.Fatalf("Fail to get section 'douban': %v", err)
	}

	Host = sec.Key("HOST").MustString("https://douban.lovemefan.top")
}

// SendRequest 发送请求
func SendRequest() {

}
