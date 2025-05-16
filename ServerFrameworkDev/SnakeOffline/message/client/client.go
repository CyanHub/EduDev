package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/gorilla/websocket"
)

func main() {
    // 连接 WebSocket 服务器
    conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8090/ws", nil)
    if err != nil {
        log.Fatal("连接失败:", err)
    }
    defer conn.Close()

    // 创建一个通道用于接收中断信号
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

    // 启动一个 goroutine 来处理接收消息
    go func() {
        for {
            _, message, err := conn.ReadMessage()
            if err != nil {
                if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                    log.Printf("读取消息错误: %v", err)
                }
                return
            }
            fmt.Printf("\n收到服务端消息: %s\n请输入要发送的消息: ", string(message))
        }
    }()

    // 处理用户输入
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("客户端可以开始发送消息（输入消息后按回车发送）:")

    for scanner.Scan() {
        message := scanner.Text()
        
        // 发送消息给服务器
        err := conn.WriteMessage(websocket.TextMessage, []byte(message))
        if err != nil {
            log.Printf("发送消息失败: %v", err)
            return
        }
        
        fmt.Println("请输入要发送的消息:")
    }
}