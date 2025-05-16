package main

import (
	"ServerFrameWork/pkg"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"sync"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 15
)

var foodX, foodY int // 食物位置

type Player struct {
	X, Y  int
	Snake [][2]float64 // 蛇的位置
	Name  string       // 玩家名称
}

var players = make(map[net.Conn]*Player) // 保存所有玩家信息（连接）
var mu sync.Mutex                        // 保护players的读写操作

func placeFood() {
	foodX = (rand.Intn(screenWidth / gridSize)) * gridSize
	foodY = (rand.Intn(screenHeight / gridSize)) * gridSize
	fmt.Printf("食物已生成, 当前位置为 x:%d, y:%d\n", foodX, foodY) // 打印食物位置

}

func main() {
	listener, err := net.Listen("tcp", ":5090")
	if err != nil {
		panic(err)
	}

	fmt.Println("服务器启动成功……")
	placeFood() // 生成食物

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("与客户端建立连接失败：", err.Error())
			continue
		}
		go HandleClient(conn) // 处理每个客户端连接
	}
}

func HandleClient(conn net.Conn) {
	defer func() {
		conn.Close()
		mu.Lock()
		delete(players, conn)
		mu.Unlock()
	}()

	// 保存玩家信息
	player := Player{}
	player.Y = 0
	player.X = 0
	player.Snake = [][2]float64{{0, 0}}
	player.Name = conn.RemoteAddr().String()
	fmt.Println("客户端连接成功", conn.RemoteAddr().String(), "开始处理消息...")
	mu.Lock()               // 加锁，防止并发读写
	players[conn] = &player // 保存玩家连接
	mu.Unlock()             // 解锁

	Broadcast(fmt.Sprintf("food#%d#%d", foodX, foodY), pkg.TypeFood)

	// 1.接收玩家信息
	for {
		messageFrame, err := pkg.ReadFrame(conn)
		// buffer := make([]byte, 2048)
		// n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("读取客户端", conn.RemoteAddr().String(), "失败：", err.Error())
			return
		}
		// i.玩家位置信息   i.玩家吃到食物信息
		// message := string(buffer[:n])
		message := string(messageFrame.Data)
		// messageSlice := strings.Split(message, "#")

		switch messageFrame.Type {
		case pkg.TypePlayerMove:
			// 收到客户端玩家信息
			newPlayer := Player{}
			err := json.Unmarshal([]byte(message), &newPlayer)
			if err != nil {
				fmt.Println("反序列化失败：", err.Error())
				continue
			}
			mu.Lock()
			players[conn] = &player
			mu.Unlock()
			// 保存之后，再广播出去
			Broadcast(message, pkg.TypePlayerMove)
			break
		case pkg.TypeEat:
			// 收到客户端吃到食物信息
			placeFood()
			Broadcast(fmt.Sprintf("food#%d#%d", foodX, foodY), pkg.TypeFood)
		}

		// if len(messageSlice) == 1 {
		// 	// 收到客户端玩家信息
		// 	newPlayer := Player{}
		// 	err := json.Unmarshal([]byte(message), &newPlayer)
		// 	if err != nil {
		// 		fmt.Println("反序列化失败：", err.Error())
		// 		continue
		// 	}
		// 	mu.Lock()
		// 	players[conn] = &player
		// 	mu.Unlock()
		// 	// 保存之后，再广播出去
		// 	Broadcast(message)
		// } else if len(messageSlice) == 2 {
		// 	// 收到客户端吃到食物信息
		// 	placeFood()
		// 	Broadcast(fmt.Sprintf("food#%d#%d", foodX, foodY))
		// }
	}
}

func Broadcast(message string, msgType uint8) {
	mu.Lock()
	defer mu.Unlock()

	for conn := range players {
		err := pkg.WriteFrame(conn, msgType, []byte(message))
		// _, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("向客户端", conn.RemoteAddr().String(), "发送消息失败：", err.Error())
			continue
		}
	}
}
