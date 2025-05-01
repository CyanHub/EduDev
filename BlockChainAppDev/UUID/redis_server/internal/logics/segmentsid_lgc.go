package logics

import (
	"fmt"
	"github.com/jiebozeng/golangutils/timeutils"
	"gorm.io/gorm"
	"log"
	"BlockChainDev/redis_server/internal/models"
	"BlockChainDev/redis_server/pkg/mysqldb"
)

type SegmentsId_lgc struct {
}

func (s *SegmentsId_lgc) GetSegmentsIds(bizType int64) (minId int64, maxId int64, err error) {
	seg := models.SegmentsId{}
	nowTime := timeutils.GetNowTime()
	query := mysqldb.Mysql.Model(&seg).Debug().Where("biz_type = ?", bizType).Limit(1)
	log.Println(query.First(&seg))
	if err = query.Find(&seg).Error; err != nil {
		return 0, 0, err
	}
	minId = seg.MaxId
	maxId = seg.MaxId + seg.Step - 1
	result := mysqldb.Mysql.Model(&seg).Debug().Where("version = ? AND biz_type = ?", seg.Version, bizType).Updates(map[string]interface{}{
		"max_id":     gorm.Expr("max_id + step"),
		"version":    gorm.Expr("version + 1"),
		"updated_at": nowTime,
	})
	fmt.Println("请求问题：", result.Error)
	if result.Error != nil {
		return 0, 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, 0, fmt.Errorf("获取ID失败，业务类型不对或者版本不对")
	}

	return minId, maxId, nil
}
