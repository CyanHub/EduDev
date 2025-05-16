package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 1.发送消息给服务器
// 2.接收服务器的消息

func main() {
	// 客户端
	conn, err := net.Dial("tcp", ":5090")
	if err != nil {
		panic(err)
	}
	fmt.Println("客户端", conn.LocalAddr().String(), "连接成功")

	go func() {
		// 接收服务器的消息
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			message := scanner.Text()
			fmt.Println(message)
		}
	}()

	// 发送消息给服务器
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		message = fmt.Sprintf("%s：%s", conn.LocalAddr().String(), message) + "\n"
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("客户端发送消息失败", err.Error())
			continue
		}
		// 移除打印自己输入消息的逻辑
	}
}
