package main

import (
    "fmt"
    "net/http"
    // "strconv"
    "time"
)

func grabRedpack(userId, redpackId int64) {
    url := fmt.Sprintf("http://127.0.0.1:8188/api/v1/redpack/grab?user_id=%d&redpack_id=%d", userId, redpackId)
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("抢红包请求失败:", err)
        return
    }
    defer resp.Body.Close()
    fmt.Printf("用户 %d 抢红包 %d 结果: %d\n", userId, redpackId, resp.StatusCode)
}

// 避免同一个包中出现多个main函数，故而将其封装并在另一个go文件中调用
func StartAutoGrab() {
    userId := int64(1)
    redpackId := int64(1)
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()
    for range ticker.C {
        grabRedpack(userId, redpackId)
    }
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func grabRedpack(userId, redpackId int64) {
// 	url := fmt.Sprintf("http://127.0.0.1:8188/api/v1/redpack/grab?user_id=%d&redpack_id=%d", userId, redpackId)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("抢红包请求失败:", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	fmt.Printf("用户 %d 抢红包 %d 结果: %d\n", userId, redpackId, resp.StatusCode)
// }

// func main() {
// 	// 模拟用户 ID
// 	userId := int64(1)
// 	// 模拟红包 ID，这里假设红包 ID 从 1 开始
// 	redpackId := int64(1)

// 	// 每隔 2 秒自动抢一次红包
// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		grabRedpack(userId, redpackId)
// 	}
// }
