package initialize

import (
	"FileSystem/global"
	// "FileSystem/model"
	"fmt"
	"log"

	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// AutoMigrate 自动迁移数据库表结构
// func AutoMigrate(db *gorm.DB) error {
// 	return db.AutoMigrate(
// 		&model.User{},
// 		&model.OperationRecord{},
// 		&model.Role{},
// 		&model.Article{},
// 		&model.Subject{},
// 	)
// }

// MustLoadGorm 初始化 GORM 数据库连接，若失败则终止程序
func MustLoadGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s?&parseTime=True&loc=Local",
		global.CONFIG.MySQL.User,
		global.CONFIG.MySQL.Password,
		global.CONFIG.MySQL.Host,
		global.CONFIG.MySQL.Port,
		global.CONFIG.MySQL.Database,
		global.CONFIG.MySQL.Config)

	// 使用 sql.Open 而不是 gorm.Open
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("SQL 数据库连接初始化失败: %v", err)
	}

	// 使用 sqlDB 创建 gorm.DB 实例
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("GORM 数据库连接初始化失败: %v", err)
	}
	global.DB = db
}
