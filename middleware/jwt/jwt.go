package jwt

import (
	"JalanJiang/douban-tool-server/pkg/e"
	"JalanJiang/douban-tool-server/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT 中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头 token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			util.ReturnError(c, http.StatusBadRequest, e.INVALID_PARAMS)
			c.Abort()
		}

		// 解析 Token
		claims, err := util.ParseToken(token)
		if err != nil {
			// Token 解析有误
			util.ReturnError(c, http.StatusUnauthorized, e.ERROR_AUTH_CHECK_TOKEN_FAIL)
			c.Abort()
		}

		// Token 过期
		if time.Now().Unix() > claims.ExpiresAt {
			util.ReturnError(c, http.StatusUnauthorized, e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT)
			c.Abort()
		}

		c.Set("userID", claims.UserID)
	}
}
