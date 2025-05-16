package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 定义一个全局 upgrader 用于将 HTTP 协议升级为 WebSocket 协议
// 定义 WebSocket 配置
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    // 允许所有来源的连接（生产环境中应该配置具体的检查逻辑）
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

// Connection 包装 websocket 连接
type Connection struct {
    conn      *websocket.Conn
    mu        sync.Mutex // 保护写操作
    closeChan chan struct{}
}

// 创建新的连接包装器
func newConnection(conn *websocket.Conn) *Connection {
    return &Connection{
        conn:      conn,
        closeChan: make(chan struct{}),
    }
}

// 安全地写入消息
func (c *Connection) writeMessage(messageType int, data []byte) error {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.conn.WriteMessage(messageType, data)
}


// 处理 WebSocket 连接的函数
func handleConnection(w http.ResponseWriter, r *http.Request) {
	 // 升级 HTTP 连接为 WebSocket
	 conn, err := upgrader.Upgrade(w, r, nil)
	 if err != nil {
		 log.Printf("升级连接失败: %v", err)
		 return
	 }
	 
	 // 创建连接包装器
	 c := newConnection(conn)
	 
	 // 确保连接最终会关闭
	 defer func() {
		 log.Println("关闭连接...")
		 conn.Close()
		 close(c.closeChan)
	 }()
 
	 // 设置连接参数
	 conn.SetReadLimit(512) // 限制消息大小
	 conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	 conn.SetPongHandler(func(string) error {
		 conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		 return nil
	 })
 
	 // 启动心跳检测
	 go func() {
		 ticker := time.NewTicker(54 * time.Second)
		 defer ticker.Stop()
 
		 for {
			 select {
			 case <-ticker.C:
				 if err := c.writeMessage(websocket.PingMessage, nil); err != nil {
					 log.Printf("发送 ping 失败: %v", err)
					 return
				 }
			 case <-c.closeChan:
				 return
			 }
		 }
	 }()
 
	 // 主消息处理循环
	 for {
		 messageType, message, err := conn.ReadMessage()
		 if err != nil {
			 if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				 log.Printf("读取错误: %v", err)
			 }
			 return
		 }
 
		 // 处理接收到的消息
		 log.Printf("接收到消息: %s", message)
 
		 // 发送响应
		 if err := c.writeMessage(messageType, message); err != nil {
			 log.Printf("发送消息失败: %v", err)
			 return
		 }
	 }
}

func main() {
	// 设置 WebSocket 连接的路由
	http.HandleFunc("/ws", handleConnection)

	// 启动 HTTP 服务器
	log.Println("WebSocket 服务器启动，监听端口 8090...")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}