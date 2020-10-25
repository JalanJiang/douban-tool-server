package util

import (
	"JalanJiang/douban-tool-server/pkg/e"
	"github.com/gin-gonic/gin"
)

// ReturnError 返回错误信息
func ReturnError(c *gin.Context, httpCode int, code string) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  e.MsgFlags[code],
	})
}
