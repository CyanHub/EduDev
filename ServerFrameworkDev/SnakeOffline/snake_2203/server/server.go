package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"ServerFramework/snake_2203/pkg"
	"sync"
)

type Player struct {
	X, Y  int
	Snake [][2]float64
	Name  string
}

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 15
)

var foodX, foodY int

var players = make(map[net.Conn]*Player)
var playerMu sync.Mutex

func main() {
	listener, err := net.Listen("tcp", ":10086")
	if err != nil {
		panic(err)
	}
	fmt.Println("服务器启动成功...")
	placeFood()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("服务器与客户端%s建立连接失败\n", conn.RemoteAddr().String())
			continue
		}
		go HandleClient(conn)
	}
}

func HandleClient(conn net.Conn) {
	defer func() {
		conn.Close()
		playerMu.Lock()
		delete(players, conn)
		playerMu.Unlock()
	}()

	//Broadcast(fmt.Sprintf("food#%d#%d", foodX, foodY))

	//BroadcastPlayerInfo(conn)

	playerMu.Lock()
	player := Player{}
	player.Name = conn.RemoteAddr().String()
	player.X = 0
	player.Y = 0
	player.Snake = [][2]float64{{0, 0}}
	players[conn] = &player
	playerMu.Unlock()

	BroadcastWithFrame(fmt.Sprintf("food#%d#%d", foodX, foodY), pkg.TypeFood)

	// 1. 接受玩家发送的消息  玩家状态  吃到食物
	for {
		messageFrame, err := pkg.ReadFrame(conn)
		if err != nil {
			fmt.Printf("读取客户端%s消息失败", conn.RemoteAddr().String())
			return
		}
		message := string(messageFrame.Data)
		switch messageFrame.Type {
		case pkg.TypePlayerMove:
			// 处理玩家信息  1. 更新服务器保存的玩家信息  2. 广播给其他客户端
			newPlayer := Player{}
			err = json.Unmarshal([]byte(message), &newPlayer)
			if err != nil {
				fmt.Println("反序列化失败, 原因是：", err.Error())
				continue
			}
			players[conn] = &newPlayer
			//Broadcast(message)
			BroadcastWithFrame(message, pkg.TypePlayerMove)
			break
		case pkg.TypeEat:
			fmt.Println("收到吃到食物的消息", message)
			// 处理玩家吃到食物的信息
			placeFood()
			//Broadcast(fmt.Sprintf("food#%d#%d", foodX, foodY))
			BroadcastWithFrame(fmt.Sprintf("food#%d#%d", foodX, foodY), pkg.TypeFood)
		}
	}
}

func placeFood() {
	foodX = (rand.Intn(screenWidth / gridSize)) * gridSize
	foodY = (rand.Intn(screenHeight / gridSize)) * gridSize
	fmt.Println("Food placed")
}

//	func Broadcast(message string) {
//		playerMu.Lock()
//		defer playerMu.Unlock()
//		for conn := range players {
//			_, err := conn.Write([]byte(message))
//			if err != nil {
//				fmt.Printf("向客户端%s发送数据失败，原因是：%s", conn.RemoteAddr().String(), err.Error())
//				continue
//			}
//		}
//	}
func BroadcastWithFrame(message string, msgType uint8) {
	playerMu.Lock()
	defer playerMu.Unlock()
	fmt.Println("当前用户数量：", len(players))
	fmt.Println(message)
	for conn := range players {
		err := pkg.WriteFrame(conn, msgType, []byte(message))
		if err != nil {
			fmt.Printf("向客户端%s发送数据失败，原因是：%s", conn.RemoteAddr().String(), err.Error())
			continue
		}
		fmt.Println("向客户端", conn.RemoteAddr().String(), "发送消息成功")
	}
}

//func BroadcastPlayerInfo(conn net.Conn) {
//	playerMu.Lock()
//	defer playerMu.Unlock()
//
//	for _, player := range players {
//		data, err := json.Marshal(&player)
//		if err != nil {
//			fmt.Println("序列化失败， 原因是：", err.Error())
//			continue
//		}
//		_, err = conn.Write(data)
//		if err != nil {
//			fmt.Println("向客户端写数据失败，原因是：", err.Error())
//		}
//	}
//}
