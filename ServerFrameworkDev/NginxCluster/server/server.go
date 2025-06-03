package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("请输入端口号")
        return
    }
    // 获取启动时输入的端口号
    port := os.Args[1]
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("接收到请求: %s %s\n", r.Method, r.URL.Path)
        fmt.Fprintf(w, "pong 来自端口: "+port)
    })
    fmt.Println("http 服务启动,监听端口： ", port)
    http.ListenAndServe(":"+port, nil)
}