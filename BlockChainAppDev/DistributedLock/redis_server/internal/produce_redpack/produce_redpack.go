package produce_redpack

import (
	"BlockChainDev/redis_server/internal/models"
	"BlockChainDev/redis_server/pkg/logs"
	"BlockChainDev/redis_server/pkg/mysqldb"
	"BlockChainDev/redis_server/pkg/redisdb"
	"context"
	"encoding/json"
	"github.com/jiebozeng/golangutils/convert"
	"github.com/jiebozeng/golangutils/timer"
	"strconv"
	"time"
)

func Init() {
	tm := timer.NewTimer(time.Second)
	tm.Start()

	tm.AddTimer(time.Second*5, 10, func(t *timer.Timer) {
		// 生成红包 插入数据库 并写入 Redis
		redpack := &models.Redpack{
			Amount:    100,
			Num:       10,
			ValidTime: -1,
			Status:    strconv.Itoa(1), // 将整数 1 转换为字符串 "1"
			ProNum:    0,
		}
		redpack.CreatedAt = time.Now()
		redpack.UpdatedAt = time.Now()

		// 插入红包信息到 MySQL
		err := mysqldb.Mysql.Create(redpack).Error
		if err != nil {
			logs.ZapLogger.Error("生成红包失败: " + err.Error())
			return
		}
		logs.ZapLogger.Info("生成红包成功 id=>" + convert.ToString(redpack.ID))

		// 红包信息保存到 Redis
		redpackMarshal, err := json.Marshal(redpack)
		if err != nil {
			logs.ZapLogger.Error("序列化红包信息失败: " + err.Error())
			return
		}
		ctx := context.Background()
		err = redisdb.RedisClient.Set(ctx, "redpack:"+convert.ToString(redpack.ID), redpackMarshal, 0).Err()
		if err != nil {
			logs.ZapLogger.Error("保存红包信息到 Redis 失败: " + err.Error())
		}
	})
}
