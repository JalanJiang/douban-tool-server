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

// Topics 话题列表
type Topics struct {
	Count      int `json:"count"`
	Start      int `json:"start"`
	TopicItems []struct {
		ID            string `json:"id"`
		UpdateTime    string `json:"update_time"`
		Title         string `json:"title"`
		CreateTime    string `json:"create_time"`
		CommentsCount int    `json:"comments_count"`
	} `json:"topics"`
	Total int `json:"total"`
}

// GetTopic 获取话题详情
func GetTopic(c *gin.Context) {

}

// GetTopics 获取话题列表
func GetTopics(c *gin.Context) {
	groupID := c.Query("group_id")
	keyWord := c.Query("key_word")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	start := (page - 1) * pageSize

	// 创建验证器
	valid := validation.Validation{}
	valid.Required(groupID, "group_id").Message("小组 ID 不能为空")
	valid.Required(keyWord, "key_word").Message("话题关键词不能为空")
	valid.Numeric(page, "page").Message("页码必须是数字")
	valid.Numeric(pageSize, "pageSize").Message("每页数量必须是数字")

	// 构造查询内容
	uri := fmt.Sprintf("/api/v2/group/%s/search/topic", groupID)
	params := fmt.Sprintf("q=%s&start=%d&count=%d&sortby=new", keyWord, start, pageSize)

	resp, err := util.SendRequest(uri, params)
	if err != nil {
		// 请求发生错误时
		log.Fatalf("Request failed: %v", err)
	}

	// JSON 解析
	var result Topics
	jsonDecodeErr := json.Unmarshal(resp, &result)
	if err != nil {
		log.Fatalf("JSON decode failed: %v", jsonDecodeErr)
	}

	responseData := make(map[string]interface{})
	responseData["total"] = result.Total
	responseData["items"] = result.TopicItems

	c.JSON(http.StatusOK, gin.H{
		"data": responseData,
		"code": 0,
		"msg":  "test",
	})
}
