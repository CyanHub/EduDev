package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"FileSystem/global"
	"FileSystem/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(db *gorm.DB) error {
	// 只迁移 file_system_fixed.sql 中定义的表对应的模型
	return db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.UserRole{}, // 添加 UserRole 结构体
		&model.RolePermission{},
		&model.File{},
		&model.FilePermission{},
		&model.OperationRecord{},
	)
}

// MustInitDB 初始化 MySQL 数据库
func MustInitDB() {
	mysqlConf := global.CONFIG.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&%s",
		mysqlConf.User,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.Database,
		mysqlConf.Config,
	)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// 配置连接池
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	global.DB = db
	if err := AutoMigrate(global.DB); err != nil {
		panic(fmt.Sprintf("数据库表结构迁移失败: %v", err))
	}
}

// MustInitRedis 初始化 Redis 客户端
func MustInitRedis() {
	redisConf := global.CONFIG.Redis
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})

	// 测试 Redis 连接
	ctx := global.Context
	if _, err := global.Redis.Ping(ctx).Result(); err != nil {
		panic(fmt.Sprintf("Redis 连接失败: %v", err))
	}
}
