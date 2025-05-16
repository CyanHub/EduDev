package gorm

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func SessionDryRun() {
	var class Class
	stmt := DB.Session(&gorm.Session{DryRun: true}).Model(&class).Where("name = ?", "一班").Find(&class).Statement
	fmt.Println(stmt.SQL.String())
	fmt.Println(stmt.Vars)
	fmt.Println(class)
}

func SessionNewDB() {
	db := DB.Where("name = ?", "一班").Session(&gorm.Session{DryRun: true})
	var class Class
	stmt := db.Find(&class).Statement
	fmt.Println(stmt.SQL.String())
	tx := db.Session(&gorm.Session{DryRun: true, NewDB: true})
	stmt = tx.Find(&class).Statement
	fmt.Println(stmt.SQL.String())
}

func SessionPrepareStmt() {
	db := DB.Model(&Class{}).Session(&gorm.Session{PrepareStmt: true})
	start := time.Now()
	for i := 0; i < 1000; i++ {
		db.Find(&Class{}, 1)
	}
	fmt.Println("开启预编译：", time.Since(start))
	stmtManager, ok := db.ConnPool.(*gorm.PreparedStmtDB)
	if !ok {
		fmt.Println("db.ConnPool is not a *gorm.PreparedStmtDB")
		return
	}
	fmt.Println(stmtManager.Stmts)
	start = time.Now()
	db2 := DB.Model(&Class{})
	for i := 0; i < 1000; i++ {
		db2.Find(&Class{}, 1)
	}
	fmt.Println("不开启预编译：", time.Since(start))
}
