package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功")

	// 启动一个 goroutine 来接收服务器的消息
	go func() {
		scanner := bufio.NewScanner(conn)
		// 持续扫描服务器发来的消息
		for scanner.Scan() {
			// 打印接收到的消息
			fmt.Println(scanner.Text())
		}
	}()

	// 主 goroutine 负责读取用户输入并发送
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// 发送消息到服务器，添加换行符确保服务器能正确分割消息
		_, err := conn.Write([]byte(scanner.Text() + "\n"))
		if err != nil {
			fmt.Println("发送失败:", err)
			return
		}
	}
}
