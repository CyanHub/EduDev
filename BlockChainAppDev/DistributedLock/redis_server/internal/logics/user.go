package logics

import (
	"BlockChainDev/redis_server/internal/models"
	"BlockChainDev/redis_server/pkg/mysqldb"
	//"context"
	"errors"
)

type User_lgc struct{}

// GetUserByUid 根据用户ID获取用户信息
func (u *User_lgc) GetUserByUid(userId int64) (*models.User, error) {
	var user models.User
	err := mysqldb.Mysql.Where("id = ?", userId).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	return &user, nil
}

type RedpackRecordLgc struct{}

// InsertRedPackRecord 插入红包记录
func (r *RedpackRecordLgc) InsertRedPackRecord(record *models.RedPackRecord) error {
	return mysqldb.Mysql.Create(record).Error
}
