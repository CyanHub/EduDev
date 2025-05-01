package router

import (
	"BlockChainDev/redis_server/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRedpackRouter 初始化红包相关路由
func InitRedpackRouter(r *gin.Engine) {
	redpackGroup := r.Group("/api/v1/redpack")
	{
		// 添加获取红包列表的路由
		redpackGroup.GET("/", v1.GetRedpacks)
		// 添加创建红包的路由
		redpackGroup.POST("/", v1.CreateRedpack)
	}
}

//package router
//
//import (
//	"github.com/gin-gonic/gin"
//)
//
//// InitRedpackRouter 初始化红包相关路由
//func InitRedpackRouter(r *gin.Engine) {
//	redpackGroup := r.Group("/api/v1/redpack")
//	{
//		// 这里可以添加具体的路由处理函数
//		// redpackGroup.GET("/", v1.GetRedpacks)
//		// redpackGroup.POST("/", v1.CreateRedpack)
//	}
//}
