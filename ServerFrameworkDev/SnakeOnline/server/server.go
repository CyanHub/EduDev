package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

// 1.保存所有客户端信息
// 2.接收客户端发送的消息
// 3.将消息广播给其他客户端

var (
	clients = make(map[string]net.Conn) // 保存所有客户端信息（连接）
	mu      sync.Mutex                  // 保护clients的读写操作
)

// Broadcast 广播消息
func Broadcast(message string) {
	mu.Lock()         // 加锁，防止并发读写
	defer mu.Unlock() // 解锁

	for _, conn := range clients { // 遍历所有客户端连接
		_, err := conn.Write([]byte(message)) // 发送消息
		if err != nil {
			fmt.Printf("客户端%s发送消息失败%s\n", conn.RemoteAddr().String(), err.Error())
			continue
		}
	}
}

// HandleClient 处理每个客户端连接
func HandleClient(conn net.Conn) {
	defer func() { // 关闭连接
		conn.Close()                                // 关闭连接
		mu.Lock()                                   // 加锁，防止并发读写
		delete(clients, conn.RemoteAddr().String()) // 删除客户端连接
		mu.Unlock()                                 // 解锁
		Broadcast(fmt.Sprintf("客户端%s断开连接\n", conn.RemoteAddr().String()) + "\n")
	}()

	fmt.Println("客户端连接成功", conn.RemoteAddr().String(), "开始处理消息...")
	// 1.保存客户端信息
	mu.Lock()                                  // 加锁，防止并发读写
	clients[conn.RemoteAddr().String()] = conn // 保存客户端连接
	mu.Unlock()                                // 解锁

	// 2.接收客户端发送的消息
	scanner := bufio.NewScanner(conn) // 读取客户端发送的消息
	for scanner.Scan() {
		message := scanner.Text() // 读取消息
		message = message + "\n"  // 添加换行符
		// 3.将消息广播给其他客户端
		Broadcast(message) // 广播消息
	}
}

func main() {
	// 修改监听地址格式
	listener, err := net.Listen("tcp", ":5090")
	if err != nil {
		panic(err)
	}
	defer listener.Close() // 关闭监听
	fmt.Println("服务器启动成功，等待客户端连接...")
	for {
		conn, err := listener.Accept() // 接收客户端连接
		if err != nil {
			fmt.Println("客户端连接失败", err.Error())
			continue // 继续等待下一个连接
		}
		go HandleClient(conn) // 处理每个客户端连接
	}
}
