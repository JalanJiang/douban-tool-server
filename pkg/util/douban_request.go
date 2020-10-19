package util

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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

	// 初始化值
	Host = sec.Key("HOST").MustString("https://douban.lovemefan.top")
	Key = os.Getenv("DOUBAN_KEY")
	APIKey = os.Getenv("DOUBAN_API_KEY")
}

// SendRequest 发送请求
func SendRequest(uri string, params string) ([]byte, error) {
	// 获取时间戳
	timestamp := time.Now().Unix()

	// 拼接 data
	data := fmt.Sprintf("GET&%s&%d", uri, timestamp)

	// 使用 key 对 data 进行 SHA1 加密
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	sig := hex.EncodeToString(sha1.Sum([]byte("")))

	// 设置请求参数
	url := fmt.Sprintf("%s%s?os_rom=android&_sig=%s&apiKey=%s&ts=%d&channel=Douban&%s", Host, uri, sig, APIKey, timestamp, params)
	req, _ := http.NewRequest("GET", url, nil)
	fmt.Println(url)

	// 设置请求头
	req.Header.Set("User-agent", "Rexxar-Core/0.1.3 api-client/1 com.douban.frodo/6.32.0(180) Android/22 product/R11 vendor/OPPO model/OPPO R11  rom/android  network/wifi  platform/mobile com.douban.frodo/6.32.0(180) Rexxar/1.2.151  platform/mobile 1.2.151")
	req.Header.Set("Host", "frodo.douban.com")
	req.Header.Set("Connection", "keep-Alive")
	req.Header.Set("Accept-Encoding", "") // gzip 返回二进制，将出现乱码

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Query topic failed: %v", err)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	return responseData, err
}
