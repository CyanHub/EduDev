package v1

import (
	"BlockChainDev/web_server/internal/logics"
	"github.com/gin-gonic/gin"
	"github.com/jiebozeng/golangutils/convert"
)

func GetUser(c *gin.Context) {
	userId := c.Param("user_id")
	userLgc := &logics.User_lgc{}
	user, err := userLgc.GetUserByUid(convert.ToInt64(userId))
	if err != nil {
		c.JSON(-1, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(0, gin.H{
		"user": user,
	})
}
