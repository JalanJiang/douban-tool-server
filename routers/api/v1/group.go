package v1

import (
	"JalanJiang/douban-tool-server/pkg/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// Groups 小组列表
type Groups struct {
	Count      int `json:"count"`
	Start      int `json:"start"`
	Total      int `json:"total"`
	GroupItems []struct {
		TopicCount          int    `json:"topic_count"`
		DescAbstract        string `json:"desc_abstract"`
		CreateTime          string `json:"create_time"`
		ID                  string `json:"id"`
		ChannelID           string `json:"channel_id"`
		Type                string `json:"type"`
		Name                string `json:"name"`
		SharingURL          string `json:"sharing_url"`
		URL                 string `json:"url"`
		URI                 string `json:"uri"`
		JoinType            string `json:"join_type"`
		Avatar              string `json:"avatar"`
		BackgroundMaskColor string `json:"background_mask_color"`
	} `json:"groups"`
}

// GetGroups 获取小组列表（根据小组名称查询）
func GetGroups(c *gin.Context) {
	groupName := c.Query("group_name")                             // 小组名称
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))           // 页码
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20")) // 每页数量
	// 计算起始位置
	start := (page - 1) * pageSize

	// 创建验证器
	valid := validation.Validation{}
	valid.Required(groupName, "group_name").Message("小组名称不能为空")
	valid.Numeric(page, "page").Message("页码必须是数字")
	valid.Numeric(pageSize, "page_size").Message("每页数量必须是数字")

	// 构造查询内容
	uri := "/api/v2/group/search/group"
	// q: 查询关键字
	// start: 起始位置
	// count: 查询总数
	params := fmt.Sprintf("q=%s&start=%d&count=%d", groupName, start, pageSize) // 拼接查询参数
	resp, err := util.SendRequest(uri, params)
	if err != nil {
		// 请求发生错误时
		log.Fatalf("Request failed: %v", err)
	}

	// 对 JSON 数据进行解析
	var result Groups
	jsonDecodeErr := json.Unmarshal(resp, &result)
	if err != nil {
		log.Fatalf("JSON decode failed: %v", jsonDecodeErr)
	}

	responseData := make(map[string]interface{})
	responseData["total"] = result.Total
	responseData["items"] = result.GroupItems

	c.JSON(http.StatusOK, gin.H{
		"data": responseData,
		"code": 0,
		"msg":  "",
	})
}
