package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"ServerFramework/snake/frame"
	"sync"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 15
)

var foodX, foodY int

type Player struct {
	X, Y  int
	Snake [][2]float64
	Name  string
}

var players = make(map[net.Conn]*Player)
var playersMu sync.Mutex

func placeFood() {
	foodX = (rand.Intn(screenWidth / gridSize)) * gridSize
	foodY = (rand.Intn(screenHeight / gridSize)) * gridSize
	fmt.Println("Food placed")
}

// 接受客户端的消息  位置大小信息   吃到食物信息
// 发送某个客户端信息到其他客户端中，发送食物的位置

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	placeFood()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("建立连接失败：", err.Error())
			continue
		}
		go HandlerPlayer(conn)
	}
}

func HandlerPlayer(conn net.Conn) {
	defer func() {
		conn.Close()
		playersMu.Lock()
		delete(players, conn)
		playersMu.Unlock()
	}()

	playersMu.Lock()
	player := Player{
		X:     0,
		Y:     0,
		Snake: [][2]float64{{0, 0}},
		Name:  conn.RemoteAddr().String(),
	}
	players[conn] = &player
	playersMu.Unlock()

	Broadcast(fmt.Sprintf("food#%d#%d", foodX, foodY), frame.TypeFood)
	fmt.Println("发送食物消息")
	BroadcastPlayerInfo(conn)

	for {
		//var buffer = make([]byte, 2048)
		//n, err := conn.Read(buffer)
		frameMessage, err := frame.ReadFrame(conn)
		if err != nil {
			fmt.Println("服务器读消息失败：", err.Error())
			return
		}
		messageStr := string(frameMessage.Data)
		playerInfo := Player{}
		_ = json.Unmarshal(frameMessage.Data, &playerInfo)
		//messageSlice := strings.Split(messageStr, "#")
		switch frameMessage.Type {
		case frame.TypePlayerMove:
			playersMu.Lock()
			players[conn].X = playerInfo.X
			players[conn].Y = playerInfo.Y
			players[conn].Snake = playerInfo.Snake
			playersMu.Unlock()

			Broadcast(messageStr, frame.TypePlayerMove)
			break
		case frame.TypeEat:
			fmt.Println("食物被吃了")
			placeFood()
			Broadcast(fmt.Sprintf("food#%d#%d", foodX, foodY), frame.TypeFood)
			break
		}
	}
}
func Broadcast(message string, msgType uint8) {
	playersMu.Lock()
	defer playersMu.Unlock()

	for conn := range players {
		_, err := frame.WriteFrame(conn, msgType, []byte(message))
		//_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("写入消息失败：", err)
		}
	}
}

func BroadcastPlayerInfo(conn net.Conn) {
	playersMu.Lock()
	defer playersMu.Unlock()
	for _, player := range players {
		data, err := json.Marshal(&player)
		if err != nil {
			fmt.Println("序列化失败：", err.Error())
		}
		//_, err = conn.Write(data)
		_, err = frame.WriteFrame(conn, frame.TypePlayerMove, data)
		if err != nil {
			fmt.Println("写入消息失败：", err)
		}
	}
}
