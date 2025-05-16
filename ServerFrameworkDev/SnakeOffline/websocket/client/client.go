package main

import (
	"fmt"
	"log"
	"github.com/gorilla/websocket"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 连接 WebSocket 服务端
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8090/ws", nil)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer conn.Close()

	// 创建一个退出信号监听器，确保程序结束时关闭连接
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		conn.Close()
		os.Exit(0)
	}()

	// 发送消息给服务端
	message := "Hello, WebSocket!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Fatal("发送消息失败:", err)
	}
	fmt.Println("发送消息:", message)

	// 接收服务器返回的消息
	_, response, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("接收消息失败:", err)
	}
	fmt.Println("接收到服务器消息:", string(response))
}