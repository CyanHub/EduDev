package redisdb

import (
	"BlockChainDev/redis_server/config"
	"BlockChainDev/redis_server/pkg/logs"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

// Init 初始化 Redis 连接
func Init() {
	addr := config.CONFIG.Redis.Host + ":" + config.CONFIG.Redis.Port
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.CONFIG.Redis.Password,
		DB:       0,
	})

	_, err := RedisClient.Ping(RedisClient.Context()).Result()
	if err != nil {
		logs.ZapLogger.Error("连接 Redis 失败: " + err.Error())
	} else {
		logs.ZapLogger.Info("成功连接到 Redis")
	}
}
