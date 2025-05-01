package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jiebozeng/golangutils/convert"
	"BlockChainDev/redis_server/internal/logics"
)

func GetSegIds(c *gin.Context) {
	bizType := c.Param("biz_type")
	segLgc := &logics.SegmentsId_lgc{}
	minId, maxId, err := segLgc.GetSegmentsIds(convert.ToInt64(bizType))
	if err != nil {
		c.JSON(-1, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(0, gin.H{
		"min_id":   minId,
		"max_id":   maxId,
		"biz_type": bizType,
	})

}
