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

// Topic 话题
type Topic struct {
	Abstract      string `json:"abstract"`
	LikeCount     int    `json:"like_count"`
	IsReportStock bool   `json:"is_report_stock"`
	ID            string `json:"id"`
	Group         struct {
		TopicCount          int           `json:"topic_count"`
		URI                 string        `json:"uri"`
		CreateTime          string        `json:"create_time"`
		ID                  string        `json:"id"`
		MemberCount         int           `json:"member_count"`
		ChannelID           string        `json:"channel_id"`
		TopicTagsEpisode    []interface{} `json:"topic_tags_episode"`
		Name                string        `json:"name"`
		URL                 string        `json:"url"`
		DescAbstract        string        `json:"desc_abstract"`
		UnreadCount         int           `json:"unread_count"`
		Avatar              string        `json:"avatar"`
		BackgroundMaskColor string        `json:"background_mask_color"`
	} `json:"group"`
	Author struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
		ID     string `json:"id"`
		UID    string `json:"uid"`
	} `json:"author"`
	Label           string      `json:"label"`
	Content         string      `json:"content"`
	GalleryTopic    interface{} `json:"gallery_topic"`
	Type            string      `json:"type"`
	IsAd            bool        `json:"is_ad"`
	CanAuthorDelete bool        `json:"can_author_delete"`
	AdFilterType    int         `json:"ad_filter_type"`
	UpdateTime      string      `json:"update_time"`
	IsElite         bool        `json:"is_elite"`
	CoverURL        string      `json:"cover_url"`
	Photos          []struct {
		Layout string `json:"layout"`
		Title  string `json:"title"`
		SeqID  string `json:"seq_id"`
		Image  struct {
			Large struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"large"`
			IsAnimated bool `json:"is_animated"`
			Normal     struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"normal"`
		} `json:"image"`
		CreationDate string `json:"creation_date"`
		TopicID      string `json:"topic_id"`
		ID           string `json:"id"`
		AuthorID     string `json:"author_id"`
		Alt          string `json:"alt"`
		Size         struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"size"`
	} `json:"photos"`
	IsPrivate        bool          `json:"is_private"`
	SharingURL       string        `json:"sharing_url"`
	IsDoubanAdAuthor bool          `json:"is_douban_ad_author"`
	URL              string        `json:"url"`
	TopicTags        []interface{} `json:"topic_tags"`
	Title            string        `json:"title"`
	URI              string        `json:"uri"`
	CreateTime       string        `json:"create_time"`
	CommentsCount    int           `json:"comments_count"`
}

// GetTopic 获取话题详情
func GetTopic(c *gin.Context) {
	topicID, _ := strconv.Atoi(c.Param("id"))

	// TODO 校验
	uri := fmt.Sprintf("/api/v2/group/topic/%d", topicID)
	resp, err := util.SendRequest(uri, "")
	if err != nil {
		// 请求发生错误时
		log.Fatalf("Request failed: %v", err)
	}

	// JSON 解析
	var result Topic
	jsonDecodeErr := json.Unmarshal(resp, &result)
	if err != nil {
		log.Fatalf("JSON decode failed: %v", jsonDecodeErr)
	}
	// var result

	// responseData := make(map[string]interface{})
	// responseData["topic"]

	c.JSON(http.StatusOK, gin.H{
		"data": result,
		"code": 0,
		"msg":  "",
	})
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
