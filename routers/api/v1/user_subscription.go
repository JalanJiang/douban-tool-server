package v1

import (
	"JalanJiang/douban-tool-server/models"
	"JalanJiang/douban-tool-server/pkg/e"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// GetSubscriptions 获取用户订阅列表
func GetSubscriptions(c *gin.Context) {
}

// AddSubscription 添加订阅
func AddSubscription(c *gin.Context) {
	// TODO: userID 来自登录态
	userID, _ := strconv.ParseUint(c.PostForm("user_id"), 10, 32)
	groupID := c.PostForm("group_id")
	groupName := c.PostForm("group_name")
	searchKeyword := c.PostForm("search_keyword")

	// 数据校验
	valid := validation.Validation{}
	valid.Required(userID, "user_id").Message("用户 ID 不能为空")
	valid.Required(groupID, "group_id").Message("小组 ID 不能为空")
	valid.Required(groupName, "group_name").Message("小组名称不能为空")
	valid.Required(searchKeyword, "search_keyword").Message("搜索关键字不能为空")

	_, userSubscriptionObj := models.AddUserSubscription(uint(userID), groupID, groupName, searchKeyword)
	// TODO 捕获错误
	// result.Error

	// TODO 返回自增 ID
	c.JSON(http.StatusOK, gin.H{
		"id":             userSubscriptionObj.ID,
		"user_id":        userSubscriptionObj.UserID,
		"group_id":       userSubscriptionObj.GroupID,
		"group_name":     userSubscriptionObj.GroupName,
		"search_keyword": userSubscriptionObj.SearchKeyword,
	})
}

// EditSubscription 编辑订阅（仅支持修改关键字）
func EditSubscription(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	searchKeyword := c.PostForm("search_keyword")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID 不能为空")
	// TODO 长度限制
	valid.Required(searchKeyword, "search_keyword").Message("搜索关键词不能为空")

	// 参数有误
	if valid.HasErrors() {
		errMessage := valid.Errors[0].Message
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  fmt.Sprintf("%s: %s", e.MsgFlags[e.INVALID_PARAMS], errMessage),
		})
	}

	// 判断数据是否存在
	result, userSubscriptionObj := models.GetUserSubscriptionByID(uint(id), uint(1))
	if result.RowsAffected == 0 {
		// 记录不存在
		c.JSON(http.StatusNotFound, gin.H{
			"code": e.NOT_FOUND,
			"msg":  e.MsgFlags[e.NOT_FOUND],
		})
	}

	userSubscriptionObj.SearchKeyword = searchKeyword
	// 更新数据
	models.Db.Save(&userSubscriptionObj)

	c.JSON(http.StatusNotFound, gin.H{
		"id":             userSubscriptionObj.ID,
		"user_id":        userSubscriptionObj.UserID,
		"group_id":       userSubscriptionObj.GroupID,
		"group_name":     userSubscriptionObj.GroupName,
		"search_keyword": userSubscriptionObj.SearchKeyword,
	})
}

// DeleteSubscription 删除订阅
func DeleteSubscription(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	// TODO 中间件获取用户 ID
	userID := 0

	// 查找资源是否存在
	if !models.ExistUserSubscriptionByID(uint(id), uint(userID)) {
		// 资源不存在
		c.JSON(http.StatusNotFound, gin.H{
			"code": e.NOT_FOUND,
			"msg":  e.MsgFlags[e.NOT_FOUND],
		})
	}

	// 删除目标
	models.DeleteUserSubscriptionByID(uint(id))

	c.String(http.StatusOK, "")
}
