package router

import (
	"github.com/gin-gonic/gin"
	v1 "BlockChainDev/redis_server/api/v1"
)

func InitSegmentsIdRouter(r *gin.Engine) {
	seg := r.Group("/segmentsid")
	{
		seg.GET("/:biz_type", v1.GetSegIds)
	}
}
