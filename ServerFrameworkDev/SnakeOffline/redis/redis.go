package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func BaseOperation() {
	// 上下文，用于控制请求超时等
	ctx := context.Background()
		// 设置键值
	err := client.Set(ctx, "username", "lisi", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取键值
	val, err := client.Get(ctx, "username").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("username:", val)

	// 删除键值
	err = client.Del(ctx, "username").Err()
	if err != nil {
		panic(err)
	}

	// 设置键值，并设置过期时间
	err = client.Set(ctx, "username", "王五", time.Second*3).Err()
	if err != nil {
		panic(err)
	}

	// 获取键值，并设置过期时间
	val, err = client.Get(ctx, "username").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("username:", val)

	time.Sleep(time.Second * 3)

	val, err = client.Get(ctx, "username").Result()
	if err != nil {
		fmt.Println("username:", err)
	}
	fmt.Println("username:", val)
}

func ListOperation() {
	ctx := context.Background()
	// 向列表中添加元素
	err := client.LPush(ctx, "idlist", 1, 2, 3).Err()
	if err != nil {
		panic(err)
	}

	// 获取列表中的元素
	vals, err := client.LRange(ctx, "idlist", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("idlist:", vals)

	// 删除列表中的元素
	err = client.LRem(ctx, "idlist", 1, 2).Err()
	if err != nil {
		panic(err)
	}

	vals, err = client.LRange(ctx, "idlist", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("idlist:", vals)

	// 取出列表中的元素
	val, err := client.RPop(ctx, "idlist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("idlist:", val)
}

func SetOperation() {
	ctx := context.Background()
	// 向集合中添加元素
	err := client.SAdd(ctx, "idset", 1, 2, 3).Err()
	if err != nil {
		panic(err)
	}

	// 获取集合中的元素
	vals, err := client.SMembers(ctx, "idset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("idset:", vals)

	// 删除集合中的元素
	err = client.SRem(ctx, "idset", 2).Err()
	if err != nil {
		panic(err)
	}

	vals, err = client.SMembers(ctx, "idset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("idset:", vals)

	// 获取集合中的元素数量
	num, err := client.SCard(ctx, "idset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("idset:", num)

	// 判断元素是否在集合中
	ok, err := client.SIsMember(ctx, "idset", 1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("idset:", ok)
}

func ZSetOperation() {
	ctx := context.Background()
	// 向有序集合中添加元素
	err := client.ZAdd(ctx, "userset", redis.Z{Score: 2, Member: "lisi"}, redis.Z{Score: 1, Member: "wangwu"}, redis.Z{Score: 3, Member: "zhaoliu"}, redis.Z{Score: 4, Member: "sunqi"}).Err()
	if err != nil {
		panic(err)
	}

	// 获取有序集合中的元素
	vals, err := client.ZRange(ctx, "userset", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userset:", vals)

	// 删除有序集合中的元素
	err = client.ZRem(ctx, "userset", "lisi").Err()
	if err != nil {
		panic(err)
	}

	vals, err = client.ZRange(ctx, "userset", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userset:", vals)

	// 获取有序集合中的元素数量
	num, err := client.ZCard(ctx, "userset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userset:", num)

	// 获取指定分数范围内的元素
	vals, err = client.ZRangeByScore(ctx, "userset", &redis.ZRangeBy{Min: "1", Max: "2"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userset:", vals)

	// 获取分数
	score, err := client.ZScore(ctx, "userset", "sunqi").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userset:", score)

	// 获取分数最高的元素
	val, err := client.ZPopMax(ctx, "userset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userset:", val)

	// 获取分数最低的元素
	val, err = client.ZPopMin(ctx, "userset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userset:", val)
}

func HashOperation() {
	ctx := context.Background()
	// 向哈希表中添加元素
	err := client.HSet(ctx, "userhash", "name", "lisi", "age", 20).Err()
	if err != nil {
		panic(err)
	}

	// 获取哈希表中的元素
	val, err := client.HGet(ctx, "userhash", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userhash:", val)

	// 获取哈希表中的所有元素
	vals, err := client.HGetAll(ctx, "userhash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userhash:", vals)

	// 删除哈希表中的元素
	err = client.HDel(ctx, "userhash", "name").Err()
	if err != nil {
		panic(err)
	}

	vals, err = client.HGetAll(ctx, "userhash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userhash:", vals)
}

func PubSubOperation() {
	ctx := context.Background()
	// 订阅频道
	sub := client.Subscribe(ctx, "room")
	defer sub.Close()

	// 发布消息
	err := client.Publish(ctx, "room", "hello").Err()
	if err != nil {
		panic(err)
	}
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("room:", msg.Payload)
}