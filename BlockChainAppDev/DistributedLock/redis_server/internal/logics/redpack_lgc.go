package logics

import (
	"BlockChainDev/redis_server/internal/models"
	"BlockChainDev/redis_server/pkg/logs"
	"BlockChainDev/redis_server/pkg/mysqldb"
	"BlockChainDev/redis_server/pkg/redisdb"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// RedpackLgc 处理红包相关逻辑的结构体
type RedpackLgc struct {
	userLogic   User_lgc
	recordLogic RedpackRecordLgc
}

// GetRedpackByID 根据红包 ID 从数据库中获取红包信息
func (r *RedpackLgc) GetRedpackByID(redpackId int64) (*models.Redpack, error) {
	var redpack models.Redpack
	err := mysqldb.Mysql.Where("id = ?", redpackId).First(&redpack).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("红包不存在")
		}
		return nil, err
	}
	return &redpack, nil
}

// GrabRedpack 实现用户抢红包的核心逻辑
func (r *RedpackLgc) GrabRedpack(userId, redpackId int64) (*models.RedPackRecord, error) {
	ctx := context.Background()
	lockKey := fmt.Sprintf("redpack_lock:%d", redpackId)
	lockTimeout := 5 * time.Second

	// 尝试获取分布式锁
	lockAcquired, err := redisdb.RedisClient.SetNX(ctx, lockKey, "1", lockTimeout).Result()
	if err != nil {
		logs.ZapLogger.Error("获取分布式锁时出错: " + err.Error())
		return nil, err
	}
	if !lockAcquired {
		return nil, errors.New("红包正在被其他用户操作，请稍后再试")
	}
	// 函数结束时释放锁
	defer redisdb.RedisClient.Del(ctx, lockKey)

	// 从 Redis 获取红包信息
	redpackJSON, err := redisdb.RedisClient.Get(ctx, fmt.Sprintf("redpack:%d", redpackId)).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			// Redis 中没有，从数据库获取
			redpack, err := r.GetRedpackByID(redpackId)
			if err != nil {
				return nil, err
			}
			redpackJSONBytes, err := json.Marshal(redpack)
			if err != nil {
				return nil, err
			}
			redpackJSON = string(redpackJSONBytes)
			redisdb.RedisClient.Set(ctx, fmt.Sprintf("redpack:%d", redpackId), redpackJSON, 0)
		} else {
			return nil, err
		}
	}

	var redpack models.Redpack
	err = json.Unmarshal([]byte(redpackJSON), &redpack)
	if err != nil {
		return nil, err
	}

	// 检查红包是否还有剩余
	if redpack.ProNum >= redpack.Num {
		return nil, errors.New("红包已抢完")
	}

	// 计算本次抢到的金额（简单平均分配，可按需修改）
	amount := redpack.Amount / (redpack.Num - redpack.ProNum)
	redpack.ProNum++
	redpack.Amount -= amount

	// 更新 Redis 中的红包信息
	updatedRedpackJSONBytes, err := json.Marshal(redpack)
	if err != nil {
		return nil, err
	}
	updatedRedpackJSON := string(updatedRedpackJSONBytes)
	if _, err := redisdb.RedisClient.Set(ctx, fmt.Sprintf("redpack:%d", redpackId), updatedRedpackJSON, 0).Result(); err != nil {
		logs.ZapLogger.Error("更新 Redis 中红包信息时出错: " + err.Error())
		return nil, err
	}

	// 更新数据库中的红包信息
	err = mysqldb.Mysql.Save(&redpack).Error
	if err != nil {
		logs.ZapLogger.Error("更新数据库中红包信息时出错: " + err.Error())
		return nil, err
	}

	// 插入红包记录
	record := &models.RedPackRecord{
		RedPackId: redpackId,
		UserId:    userId,
		Amount:    int64(amount),
	}
	err = r.recordLogic.InsertRedPackRecord(record)
	if err != nil {
		logs.ZapLogger.Error("插入红包记录时出错: " + err.Error())
		return nil, err
	}

	return record, nil
}

//package logics
//
//import (
//	"BlockChainDev/redis_server/internal/models"
//	"context"
//	"errors"
//	"fmt"
//	"github.com/go-redis/redis/v8"
//	"time"
//)
//
//// 抢红包 返回红包id 抢到的金额
//// 没抢到的话 返回的红包id为-1
//func (r *RedPackLgc) GradRedPack(userId int64, redpackId int64) (redId int64, amount int64, err error) {
//	ctx := context.Background()
//	lockKey := fmt.Sprintf("lock:redpack:%d", redpackId)
//	lockTimeout := 2 * time.Second
//
//	// 尝试获取分布式锁
//	lockAcquired, err := redisdb.RedisDb.SetNX(ctx, lockKey, "1", lockTimeout).Result()
//	if err != nil {
//		return -1, 0, err
//	}
//	if !lockAcquired {
//		// 未能获取锁，可能红包正在被其他用户处理
//		return -1, 0, errors.New("红包正在被其他用户处理，请稍后再试")
//	}
//	defer redisdb.RedisDb.Del(ctx, lockKey) // 确保锁被释放
//
//	// 获取红包信息
//	redpackJSON, err := redisdb.RedisDb.Get(ctx, fmt.Sprintf("redpack:%d", redpackId)).Result()
//	if err != nil {
//		if err == redis.Nil {
//			return -1, 0, errors.New("红包不存在或已过期")
//		}
//		return -1, 0, err
//	}
//
//	redpack := &models.Redpack{}
//	if err := json.Unmarshal([]byte(redpackJSON), redpack); err != nil {
//		return -1, 0, err
//	}
//
//	// 检查红包状态
//	if redpack.Status != 1 || redpack.Num <= redpack.ProNum {
//		return -1, 0, errors.New("红包已抢完或无效")
//	}
//
//	// 计算红包金额（这里简化为总金额/20，可以根据需求调整）
//	if redpack.Amount <= 5 {
//		amount = redpack.Amount
//	} else {
//		amount = redpack.Amount / 20
//		if amount <= 0 {
//			amount = 1 // 最小金额为1
//		}
//	}
//
//	// 更新已抢数量和剩余金额
//	redpack.ProNum++
//	redpack.Amount -= amount
//	if redpack.Amount == 0 {
//		redpack.Status = 5 //已领完
//		r.UpdateRedPackStatus(redpackId, 5)
//	}
//	updatedRedpackJSON, err := json.Marshal(redpack)
//	if err != nil {
//		return -1, 0, err
//	}
//
//	if _, err := redisdb.RedisDb.Set(ctx, fmt.Sprintf("redpack:%d", redpackId), updatedRedpackJSON, 0).Result(); err != nil {
//		return -1, 0, err
//	}
//	//抢到红包 插入记录
//	userLgc := &User_lgc{}
//	user, _ := userLgc.GetUserByUid(userId)
//	record := &models.RedPackRecord{
//		UserId:    int64(user.ID),
//		UserName:  user.UserName,
//		RedPackId: redpackId,
//		Amount:    amount,
//	}
//	record.CreatedAt = time.Now()
//	record.UpdatedAt = time.Now()
//	redpackRecordLgc := &RedpackRecordLgc{}
//	redpackRecordLgc.InsertRedPackRecord(record)
//	return redpackId, amount, nil
//}
