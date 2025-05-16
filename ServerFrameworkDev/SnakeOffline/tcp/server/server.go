package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听指定端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("启动失败:", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务端启动成功")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接失败:", err)
			continue
		}
		fmt.Println("新客户端连接成功")

		// 处理客户端连接
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		// 读取数据
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("接收失败:", err)
			return
		}
		fmt.Printf("收到客户端消息: %s", string(buffer[:n]))

		// 回写数据
		_, err = conn.Write([]byte("收到消息\n"))
		if err != nil {
			fmt.Println("发送失败:", err)
			return
		}
	}
}