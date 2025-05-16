package main

import (
    "bufio"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/websocket"
)

// WebSocket 配置
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

// 处理 WebSocket 连接
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    // 升级 HTTP 连接为 WebSocket 连接
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("连接升级失败: %v", err)
        return
    }
    defer conn.Close()

    log.Println("新的客户端连接established")

    // 启动一个 goroutine 来处理服务端发送消息
    go handleServerSend(conn)

    // 主循环处理接收到的消息
    for {
        _, message, err := conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                log.Printf("读取消息错误: %v", err)
            }
            return
        }
        
        // 打印接收到的消息
        log.Printf("收到客户端消息: %s", message)
    }
}

// 处理服务端发送消息
func handleServerSend(conn *websocket.Conn) {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("服务端可以开始发送消息（输入消息后按回车发送）:")

    for scanner.Scan() {
        message := scanner.Text()
        
        // 发送消息给客户端
        err := conn.WriteMessage(websocket.TextMessage, []byte(message))
        if err != nil {
            log.Printf("发送消息失败: %v", err)
            return
        }
        
        fmt.Println("请输入要发送的消息:")
    }
}

func main() {
    // 设置 WebSocket 路由
    http.HandleFunc("/ws", handleWebSocket)

    // 启动服务器
    serverAddr := ":8090"
    fmt.Printf("WebSocket 服务器启动在 %s\n", serverAddr)
    log.Fatal(http.ListenAndServe(serverAddr, nil))
}