package gorm

import (
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID        uint64         `gorm:"primaryKey" json:"id"` // 主键ID
	CreatedAt time.Time      `json:"createdAt"`            // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`            // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
	Username  string         `json:"userName" gorm:"index;comment:用户登录名"`
	Password  string         `json:"-"  gorm:"comment:用户登录密码"`
	NickName  string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	HeaderImg string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	RoleId    uint64         `json:"roleId" gorm:"default:888;comment:用户角色Id"`
	Phone     string         `json:"phone"  gorm:"comment:用户手机号"`
	Email     string         `json:"email"  gorm:"comment:用户邮箱"`
	Enable    int8           `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
	Balance   float64        `json:"balance" gorm:"type:float;comment:用户余额"`
}

func (User) TableName() string {
	return "user"
}

var DB *gorm.DB

func init() {
	_ = logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // 设置输出到标准输出
        logger.Config{
            SlowThreshold:             200 * time.Millisecond, // 慢查询阈值
            LogLevel:                  logger.Info,            // 日志级别
            IgnoreRecordNotFoundError: true,                   // 忽略 ErrRecordNotFound 错误
            Colorful:                  true,                   // 是否启用彩色日志
		},
	)
	db, err := gorm.Open(mysql.Open("root:292378@tcp(localhost:3306)/go_shop?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: nil,
	})
	if err != nil {
		panic(err)
	}
	DB = db
}

func DemoWithTransaction() {
	tx := DB.Begin()
	// 转账
	// update user set balance = balance - 100 where id = 1
	err := tx.Table("user").Where("id = ?", 1).Update("balance", gorm.Expr("balance - ?", 100)).Error
	if err != nil {
		tx.Rollback()
		return
	}
	// update user set balance = balance + 100 where id = 2
	err = tx.Table("user").Where("id = ?", 2).Update("balance", gorm.Expr("balance + ?", 100)).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}


func DemoWithTransactionError() {
	tx := DB.Begin()
	// 转账
	// update user set balance = balance - 100 where id = 1
	err := tx.Table("user").Where("id = ?", 1).Update("balance", gorm.Expr("balance - ?", 100)).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Table("user").Create(&User{
		Username: "kutqrx49",
		Password: "123456",
		Balance:  100,
	}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	// update user set balance = balance + 100 where id = 2
	err = tx.Table("user").Where("id = ?", 2).Update("balance", gorm.Expr("balance + ?", 100)).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}

func DemoWithTransactionCallback() {
	DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Table("user").Where("id = ?", 1).Update("balance", gorm.Expr("balance - ?", 100)).Error
		if err != nil {
			return err
		}
		err = tx.Table("user").Create(&User{
			Username: "kutqrx49",
			Password: "123456",
			Balance:  100,
		}).Error
		if err != nil {
			return err
		}
		// update user set balance = balance + 100 where id = 2
		err = tx.Table("user").Where("id = ?", 2).Update("balance", gorm.Expr("balance + ?", 100)).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func DemoWithSavePoint() {
	DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Table("user").Where("username = ?", "kutqrx99").Update("balance", gorm.Expr("balance + ?", 100)).Error
		if err != nil {
			return err
		}
		tx.SavePoint("step1")
		result := tx.Table("user").Where("username = ?", "kutqrx999").Update("balance", gorm.Expr("balance - ?", 100))
		if result.Error != nil || result.RowsAffected == 0 {
			tx.RollbackTo("step1")
			err = tx.Table("user").Where("username = ?", "kutqrx47").Update("balance", gorm.Expr("balance - ?", 100)).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}


func DemoWithNestedTransaction() {
	DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Table("user").Where("username = ?", "kutqrx99").Update("balance", gorm.Expr("balance + ?", 100)).Error
		if err != nil {
			return err
		}

		err = tx.Transaction(func(tx *gorm.DB) error {
			result := tx.Table("user").Where("username = ?", "kutqrx999").Update("balance", gorm.Expr("balance - ?", 100))
			if result.Error != nil || result.RowsAffected == 0 {
				return errors.New("kutqrx999 balance not enough")
			}
			return nil
		})
		if err != nil {
			err = tx.Transaction(func(tx *gorm.DB) error {
				result := tx.Table("user").Where("username = ?", "kutqrx47").Update("balance", gorm.Expr("balance - ?", 100))
				if result.Error != nil || result.RowsAffected == 0 {
					return errors.New("kutqrx47 balance not enough")
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}