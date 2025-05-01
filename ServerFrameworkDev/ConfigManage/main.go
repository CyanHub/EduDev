package main

import (
	"ConfigManage/global"
	"ConfigManage/initialize"
	"fmt"
)

func main() {
    initialize.MustConfig()
    fmt.Println(global.CONFIG)  // 打印出`global.CONFIG`变量
}
