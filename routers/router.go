package routers

import (
	"JalanJiang/douban-tool-server/middleware/jwt"
	v1 "JalanJiang/douban-tool-server/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
	r := gin.New()

	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// gin.SetMode(setting.RunMode)
	apiv1.POST("/users/login", v1.Login)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT()) // 引入中间件
	{
		// 订阅
		apiv1.GET("/subscriptions", v1.GetSubscriptions)
		apiv1.POST("/subscriptions", v1.AddSubscription)
		apiv1.PUT("/subscriptions/:id", v1.EditSubscription)
		apiv1.DELETE("/subscriptions/:id", v1.DeleteSubscription)

		// 小组
		apiv1.GET("/groups", v1.GetGroups)

		// 话题
		apiv1.GET("/topics/:id", v1.GetTopic)
		apiv1.GET("/topics", v1.GetTopics)
	}

	return r
}
