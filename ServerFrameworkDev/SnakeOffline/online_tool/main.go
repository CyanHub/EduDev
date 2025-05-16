// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

type Message struct {
	Type    string `json:"type"`    // "UPDATE" 或 "INIT"
	Content string `json:"content"` // 整个文档内容
}

var (
	clients   = make(map[*Client]bool)
	clientsMu sync.Mutex
	content   = ""
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}

	client := &Client{
		conn: conn,
		send: make(chan []byte, 256),
	}

	// 注册客户端
	clientsMu.Lock()
	clients[client] = true
	clientsMu.Unlock()

	// 发送初始内容
	initMsg, _ := json.Marshal(Message{
		Type:    "INIT",
		Content: content,
	})
	client.send <- initMsg

	go client.writePump()
	go client.readPump()
}

func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
		clientsMu.Lock()
		delete(clients, c)
		clientsMu.Unlock()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			continue
		}

		// 更新内容并广播
		content = msg.Content
		fmt.Println(content)
		broadcast(message, c)
	}
}

func (c *Client) writePump() {
	defer c.conn.Close()
	for message := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			return
		}
	}
}

func broadcast(message []byte, exclude *Client) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for client := range clients {
		if client != exclude {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
