package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	// Cfg 配置读取
	//Cfg *ini.File
	//// Host 域名
	//Host string
	// AppID 小程序应用 ID
	AppID string
	// AppSecret 小程序应用 Secret
	AppSecret string
)

func init() {
	//var err error
	//Cfg, err = ini.Load("conf/app.ini")
	//if err != nil {
	//	log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	//}

	sec, err := Cfg.GetSection("wechat")
	if err != nil {
		log.Fatalf("Fail to get section 'douban': %v", err)
	}

	// 初始化
	Host = sec.Key("HOST").MustString("https://api.weixin.qq.com")
	AppID = os.Getenv("WX_APP_ID")
	AppSecret = os.Getenv("WX_APP_SECRET")
}

// WxLogin 发送登录请求
func WxLogin(code string) ([]byte, error) {
	// GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	url := fmt.Sprintf(
		"%s/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		Host,
		AppID,
		AppSecret,
		code)
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept-Encoding", "") // gzip 返回二进制，将出现乱码

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Query topic failed: %v", err)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)

	return responseData, err
}
