package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	// 存储所有连接的客户端，key 是连接，value 是客户端地址
	clients = make(map[net.Conn]string)
	// 保护 clients 映射的互斥锁
	clientsMu sync.Mutex
)

// 处理单个客户端连接
func handleClient(conn net.Conn) {
	defer func() {
		removeClient(conn) // 客户端断开时移除
		conn.Close()       // 关闭连接
	}()

	// 将新客户端添加到映射中
	clientsMu.Lock()
	clients[conn] = conn.RemoteAddr().String()
	clientsMu.Unlock()
	fmt.Printf("客户端 %s 连接成功\n", conn.RemoteAddr())

	// 广播新客户端加入的消息
	broadcast(fmt.Sprintf("%s 加入聊天\n", conn.RemoteAddr()))

	// 持续读取客户端消息
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// 广播客户端发送的消息
		message := fmt.Sprintf("%s: %s\n", conn.RemoteAddr(), scanner.Text())
		broadcast(message)
	}
}

// 向所有客户端广播消息
func broadcast(message string) {
	// 复制当前的客户端列表，避免长时间持有锁
	clientsMu.Lock()
	currentClients := make([]net.Conn, 0, len(clients))
	for conn := range clients {
		currentClients = append(currentClients, conn)
	}
	clientsMu.Unlock()

	// 发送消息时不持有锁
	for _, conn := range currentClients {
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("发送消息失败: %v\n", err)
		}
	}
}

// 从映射中移除客户端并广播离开消息
func removeClient(conn net.Conn) {
	clientsMu.Lock()
	addr := clients[conn]
	delete(clients, conn)
	clientsMu.Unlock()

	// 在删除客户端后广播消息
	broadcast(fmt.Sprintf("%s 离开聊天\n", addr))
}

func main() {
	// 创建 TCP 服务器监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Chat server started on port 8080...")

	// 持续接受新的客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// 为每个客户端启动一个处理 goroutine
		go handleClient(conn)
	}
}
