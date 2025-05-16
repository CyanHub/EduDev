package initialize

import (
	"ServerFramework/global"
	"ServerFramework/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.OperationRecord{},
		&model.Role{},
		&model.Article{},
		&model.Subject{},
	)
}

// MustLoadGorm 初始化 GORM 数据库连接，若失败则终止程序
func MustLoadGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", global.CONFIG.MySQL.User, global.CONFIG.MySQL.Password, global.CONFIG.MySQL.Host, global.CONFIG.MySQL.Port, global.CONFIG.MySQL.Database, global.CONFIG.MySQL.Config)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		// 使用 log.Fatal 记录错误信息并终止程序
		log.Fatalf("GORM 数据库连接初始化失败: %v", err)
	}
	global.DB = db
}
