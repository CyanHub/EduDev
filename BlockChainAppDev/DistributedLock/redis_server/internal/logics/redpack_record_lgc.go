package logics

import (
	"BlockChainDev/redis_server/internal/models"
	"BlockChainDev/redis_server/pkg/mysqldb"
)

// logics/redpack_lgc.go
// GetAllRedpacks 获取所有红包列表
func (r *RedpackLgc) GetAllRedpacks() ([]models.Redpack, error) {
	var redpacks []models.Redpack
	err := mysqldb.Mysql.Find(&redpacks).Error
	if err != nil {
		return nil, err
	}
	return redpacks, nil
}

// CreateRedpack 创建红包
func (r *RedpackLgc) CreateRedpack(amount, num int) (*models.Redpack, error) {
	redpack := &models.Redpack{
		Amount: amount,
		Num:    num,
		// 其他字段可根据需要初始化
	}
	err := mysqldb.Mysql.Create(redpack).Error
	if err != nil {
		return nil, err
	}
	return redpack, nil
}

// logics/redpack_record_lgc.go
// GetRedpackRecordsByUserId 根据用户ID获取红包记录
func (r *RedpackRecordLgc) GetRedpackRecordsByUserId(userId int64) ([]models.RedPackRecord, error) {
	var records []models.RedPackRecord
	err := mysqldb.Mysql.Where("user_id = ?", userId).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}
