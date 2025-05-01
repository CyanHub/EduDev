package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    // 创建红包测试逻辑
    url := "http://127.0.0.1:8188/api/v1/redpack/"
    data := []byte(`{"amount": 100, "num": 10}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
    if err != nil {
        fmt.Println("创建请求失败:", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("发送请求失败:", err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("读取响应内容失败:", err)
        return
    }
    fmt.Println("响应状态码:", resp.StatusCode)
    fmt.Println("响应内容:", string(body))

    // 调用自动抢红包逻辑
    StartAutoGrab()
}

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	url := "http://127.0.0.1:8188/api/v1/redpack/"
// 	// 要发送的 JSON 数据
// 	data := []byte(`{"amount": 100, "num": 10}`)

// 	// 创建一个新的 HTTP 请求
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
// 	if err != nil {
// 		fmt.Println("创建请求失败:", err)
// 		return
// 	}

// 	// 设置请求头
// 	req.Header.Set("Content-Type", "application/json")

// 	// 发送请求
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("发送请求失败:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// 读取响应内容
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("读取响应内容失败:", err)
// 		return
// 	}

// 	fmt.Println("响应状态码:", resp.StatusCode)
// 	fmt.Println("响应内容:", string(body))
// }
