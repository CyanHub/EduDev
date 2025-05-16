package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接到服务端
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功")

	// 从标准输入读取数据
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()

		// 发送数据
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("发送失败:", err)
			return
		}

		// 接收服务端回写的数据
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("接收失败:", err)
			return
		}
		fmt.Printf("服务端回复: %s", string(buffer[:n]))
	}
}