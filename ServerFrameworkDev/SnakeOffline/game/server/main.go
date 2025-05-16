package main

import (
	"fmt"
	"net"
	"ServerFramework/game/common"
	"sync"

	"golang.org/x/exp/rand"
)

type Player struct {
	Conn  net.Conn
	X, Y  int
	Name  string
	Score int
}

var (
	players      = make(map[net.Conn]*Player)
	playersMutex sync.RWMutex
	gameOver     bool
	foodX, foodY int
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 15
)

func handlePlayer(conn net.Conn) {
	sendFrame(conn, common.TypeFood, []byte(fmt.Sprintf("%d#%d", foodX, foodY)))
	defer conn.Close()
	player := &Player{Conn: conn, X: 0, Y: 0, Name: conn.RemoteAddr().String(), Score: 0}
	playersMutex.Lock()
	players[conn] = player
	playersMutex.Unlock()

	for {
		frame, err := common.ReadFrame(conn)
        if err != nil {
            fmt.Println("Error reading frame:", err)
            break
        }

        switch frame.Type {
        case common.TypeEat:
            broadcastFrame(common.TypeEat, frame.Data)
            placeFood()
            foodData := []byte(fmt.Sprintf("%d#%d", foodX, foodY))
            broadcastFrame(common.TypeFood, foodData)
        case common.TypePlayerMove:
            broadcastFrame(common.TypePlayerMove, frame.Data)
        }
	}

	playersMutex.Lock()
	delete(players, conn)
	playersMutex.Unlock()
	// broadcast(fmt.Sprintf("%s left the game.", player.Name))
	broadcastFrame(common.TypePlayerLeft, []byte(fmt.Sprintf("%s left the game.", player.Name)))
}

func sendFrame(conn net.Conn, msgType uint8, data []byte) {
	err := common.WriteFrame(conn, msgType, data)
    if err != nil {
        fmt.Println("Error sending frame:", err)
    }
}

func broadcastFrame(msgType uint8, data []byte) {
	playersMutex.RLock()
    defer playersMutex.RUnlock()
    for _, player := range players {
        common.WriteFrame(player.Conn, msgType, data)
    }
}

func placeFood() {
	foodX = (rand.Intn(screenWidth / gridSize)) * gridSize
	foodY = (rand.Intn(screenHeight / gridSize)) * gridSize
	fmt.Println("Food placed")
}


func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Game server started on :8080")

	placeFood()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		go handlePlayer(conn)
	}
}
