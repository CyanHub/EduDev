package main

import (
	// "ServerFramework/global"
	// "ServerFramework/initialize"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// initialize.MustConfig()           // 初始化配置文件
	// initialize.MustInitRedis()        // 初始化Redis
	// initialize.MustLoadZap()          // 初始化日志
	// initialize.RegisterSerializer()   // 注册序列化器
	// initialize.MustLoadGorm()         // 初始化数据库
	// initialize.AutoMigrate(global.DB) // 自动迁移数据库
	// initialize.MustCasbin()           // 初始化Casbin
	// initialize.MustRunWindowServer()  // 初始化窗口服务器
	// StringOption()                    // 测试Redis01
	// ListOperation()                   // 测试Redis02
    // SetOperation()                    // 测试Redis03
    // HashOperation()                   // 测试Redis04
    ZSetOperation()                   // 测试Redis05
}

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 替换为 Redis 服务器地址和端口
		Username: "",               // 如果 Redis 服务器需要用户名认证，替换为实际用户名
		Password: "",               // 如果 Redis 服务器需要密码认证，替换为实际密码
		DB:       0,                // Redis 数据库编号，默认为 0,

		// 下面是一些可选的配置项，根据实际需求进行设置
		// Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		// Password: redisConf.Password,
		// DB:       redisConf.DB,
	})
}

// StringOption 测试Redis的字符串操作
func StringOption() {
	// 上下文，用于控制请求超时等
	ctx := context.Background()
	// 设置键值
	err := client.Set(ctx, "username", "洛青", 0).Err()
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
	err = client.Set(ctx, "username", "伊鸠", time.Second*3).Err()
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

// ListOperation 测试Redis的列表操作
// 列表是一个有序的字符串集合，它可以存储多个字符串元素，并且可以根据索引访问和操作元素。
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
    // vals 是一个字符串切片，包含列表中的所有元素
    // 这里是会输出 ["3", "2", "1"]，因为列表是先进后出的
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
    // 这里会输出3和1
	fmt.Println("idlist:", vals)

	// 取出列表中的元素
	val, err := client.RPop(ctx, "idlist").Result()
	if err != nil {
		panic(err)
	}
    // 这里是会输出 "3"，因为列表中只剩下一个元素了，所以取出的元素就是 "3"
	fmt.Println("idlist:", val)
}

// SetOperation 测试Redis的集合操作
// 集合是一个无序的字符串集合，它可以存储多个字符串元素，并且可以根据元素的值进行操作。
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
    // 这里是会输出 [1 2 3]，因为集合是无序的，所以取出的元素的顺序是不确定的
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
    // 这里是会输出 [1 3]，因为集合中只剩下一个元素了，所以取出的元素就是 "3"
	fmt.Println("idset:", vals)

	// 获取集合中的元素数量
	num, err := client.SCard(ctx, "idset").Result()
	if err != nil {
		panic(err)
	}
    // 这里是会输出 2
	fmt.Println("idset:", num)

	// 判断元素是否在集合中
	ok, err := client.SIsMember(ctx, "idset", 1).Result()
	if err != nil {
		panic(err)
	}
    // 这里是会输出 true
	fmt.Println("idset:", ok)

}

// ZSetOperation 测试Redis的有序集合操作
// 有序集合是一个有序的字符串集合，它可以存储多个字符串元素，并且可以根据元素的值进行操作。
func ZSetOperation() {
	ctx := context.Background()
	// 向有序集合中添加元素
	err := client.ZAdd(ctx, "userset", redis.Z{Score: 2, Member: "洛青"}, redis.Z{Score: 1, Member: "伊鸠"}, redis.Z{Score: 3, Member: "宫常"}, redis.Z{Score: 4, Member: "皓音"}).Err()
	if err != nil {
		panic(err)
	}

	// 获取有序集合中的元素
	vals, err := client.ZRange(ctx, "userset", 0, -1).Result()
	if err != nil {
		panic(err)
	}
    // 这里是会输出 ["伊鸠", "洛青", "宫常", "皓音"]，因为有序集合是有序的，所以取出的元素的顺序是确定的
	fmt.Println("userset:", vals)

	// 删除有序集合中的元素
	err = client.ZRem(ctx, "userset", "洛青").Err()
	if err != nil {
		panic(err)
	}

	vals, err = client.ZRange(ctx, "userset", 0, -1).Result()
	if err != nil {
		panic(err)
	}
    // 这里是会输出 ["伊鸠", "宫常", "皓音"]
	fmt.Println("userset:", vals)

	// 获取有序集合中的元素数量
	num, err := client.ZCard(ctx, "userset").Result()
	if err != nil {
		panic(err)
	}
    // 这里是会输出 3
	fmt.Println("userset:", num)

	// 获取指定分数范围内的元素
	vals, err = client.ZRangeByScore(ctx, "userset", &redis.ZRangeBy{Min: "1", Max: "2"}).Result()
	if err != nil {
		panic(err)
	}
    // 这里是会输出 ["伊鸠"]
	fmt.Println("userset:", vals)

	// 获取分数
	score, err := client.ZScore(ctx, "userset", "皓音").Result()
	if err != nil {
		panic(err)
	}
    // 这里是会输出 4
	fmt.Println("userset:", score)

	// 获取分数最高的元素
	val, err := client.ZPopMax(ctx, "userset").Result()
	if err != nil {
		panic(err)
	}
    // 这里是会输出 ["4","皓音" ]，因为有序集合是有序的，所以取出的元素的顺序是确定的
	fmt.Println("userset:", val)

	// 获取分数最低的元素
	val, err = client.ZPopMin(ctx, "userset").Result()
	if err != nil {
		panic(err)
	}
    // 这里是会输出 ["1","伊鸠" ]，因为有序集合是有序的，所以取出的元素的顺序是确定的
	fmt.Println("userset:", val)
}