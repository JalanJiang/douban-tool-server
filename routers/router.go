package routers

import (
	v1 "JalanJiang/douban-tool-server/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
	r := gin.New()

	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// gin.SetMode(setting.RunMode)

	// TODO: 中间件

	apiv1 := r.Group("/api/v1")
	{
		// 用户
		apiv1.POST("/users/login", v1.Login)

		// 订阅
		apiv1.GET("/subscriptions", v1.GetSubscriptions)
		apiv1.POST("/subscriptions", v1.AddSubscription)
		apiv1.PUT("/subscriptions/:id", v1.EditSubscription)
		apiv1.DELETE("/subscriptions/:id", v1.DeleteSubscription)

		// 小组
		apiv1.GET("/groups", v1.GetGroups)

		// 话题
		apiv1.GET("/topics/:id", v1.GetTopic)
	}

	return r
}
