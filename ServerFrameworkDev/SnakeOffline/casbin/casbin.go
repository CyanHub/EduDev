package casbin

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CasbinWithFile(modelPath, policyPath string) {
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		panic(err)
	}
	allow, err := enforcer.Enforce("alice", "data1", "read")
	if err != nil {
		panic(err)
	}
	if allow {
		fmt.Println("alice允许读取data1")
	} else {
		fmt.Println("alice拒绝读取data1")
	}
}

func CasbinDynamic(modelPath, policyPath string) {
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		panic(err)
	}
	allow, err := enforcer.Enforce("bob", "data2", "write")
	if err != nil {
		panic(err)
	}
	if allow {
		fmt.Println("bob允许写入data2")
	} else {
		fmt.Println("bob拒绝写入data2")
	}
	enforcer.AddPolicy("bob", "data2", "write")
	allow, err = enforcer.Enforce("bob", "data2", "write")
	if err != nil {
		panic(err)
	}
	if allow {
		fmt.Println("bob允许写入data2")
	} else {
		fmt.Println("bob拒绝写入data2")
	}
	allow, err = enforcer.Enforce("jack", "data1", "write")
	if err != nil {
		panic(err)
	}
	if allow {
		fmt.Println("jack允许写入data1")
	} else {
		fmt.Println("jack拒绝写入data1")
	}
}

func CasbinSavePolicy(modelPath, policyPath string) {
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		panic(err)
	}
	enforcer.AddPolicy("bob", "data2", "write")
	enforcer.SavePolicy()
	fmt.Println("策略已保存")
}

func CasbinWithAdapter(modelPath, policyPath string) {
	db, err := gorm.Open(mysql.Open("root:292378@tcp(localhost:3306)/ServerFrameworkelop?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	adater, err := adapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}

	enforcer, err := casbin.NewEnforcer(modelPath, adater)
	if err != nil {
		panic(err)
	}
	enforcer.AddPolicy("bob", "data2", "write")
	fmt.Println("策略已保存")
}

