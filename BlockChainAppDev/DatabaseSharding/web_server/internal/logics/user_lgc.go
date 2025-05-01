package logics

import (
	"BlockChainDev/web_server/internal/models"
	"BlockChainDev/web_server/pkg/mysqldb"
	"github.com/jiebozeng/golangutils/convert"
)

type User_lgc struct {
}

func (u *User_lgc) GetUserByUid(userId int64) (*models.User, error) {
	user := &models.User{}
	//关键点 表名不是固定的了，是根据用户id取模
	tableName := "user_" + convert.ToString(userId%10+1)
	query := mysqldb.Mysql.Table(tableName).Model(user)
	//查询单个用户
	query = query.Where("user_id = ?", userId)
	err := query.Find(&user).Limit(1).Error
	return user, err
}
