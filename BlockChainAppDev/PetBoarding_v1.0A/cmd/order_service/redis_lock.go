package main

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisLock struct {
	client *redis.Client
	key    string
	token  string
}

func NewRedisLock(client *redis.Client, key string) *RedisLock {
	return &RedisLock{
		client: client,
		key:    key,
		token:  generateToken(),
	}
}

func (l *RedisLock) Lock(ctx context.Context, ttl time.Duration) (bool, error) {
	return l.client.SetNX(ctx, l.key, l.token, ttl).Result()
}

func (l *RedisLock) Unlock(ctx context.Context) error {
	script := redis.NewScript(`
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
	`)

	result, err := script.Run(ctx, l.client, []string{l.key}, l.token).Int64()
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("解锁失败: 令牌不匹配或密钥不存在")
	}
	return nil
}

func generateToken() string {
	return "unique_token_" + time.Now().Format(time.RFC3339Nano)
}
