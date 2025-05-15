package initialize

import (
	"ServerFramework/global"
	"ServerFramework/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.OperationRecord{},
		&model.Role{},
		&model.Article{},
		&model.Subject{},
	)

}

func MustLoadGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", global.CONFIG.MySQL.User, global.CONFIG.MySQL.Password, global.CONFIG.MySQL.Host, global.CONFIG.MySQL.Port, global.CONFIG.MySQL.Database, global.CONFIG.MySQL.Config)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		err.Error()
	}
	global.DB = db
}
