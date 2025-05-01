package initialize

import (
	"ServerFramework/global"
	"ServerFramework/model"
	
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Role{},
	)
}

func MustLoadGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", global.CONFIG.MySQL.Username, global.CONFIG.MySQL.Password,
		global.CONFIG.MySQL.Host, global.CONFIG.MySQL.Port, global.CONFIG.MySQL.Database)
		fmt.Println("我的mysql",dsn)
	//dsn := "root:123456@tcp(127.0.0.1:3306)/go_shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	global.DB = db
}
