package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"ServerFramework/global"
	"ServerFramework/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.OperationRecord{},
		&model.Subject{},
	)
}

func MustInitDB() {
	sqlDB, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		panic(err)
	}
	global.DB = db
}

func MustInitRedis() {
	redisConf := global.CONFIG.Redis
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})
}
