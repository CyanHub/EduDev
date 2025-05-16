package main

import (
	"fmt"
	"net"
	"strings"
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

func broadcast(message string) {
	playersMutex.RLock()
	defer playersMutex.RUnlock()
	for _, player := range players {
		_, err := fmt.Fprintln(player.Conn, message)
		if err != nil {
			fmt.Println("Error broadcasting:", err)
		}
	}
}

func sendMessage(conn net.Conn, message string) {
	_, err := fmt.Fprintln(conn, message)
	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}

func handlePlayer(conn net.Conn) {
	sendMessage(conn, fmt.Sprintf("food#%d#%d", foodX, foodY))
	defer conn.Close()
	player := &Player{Conn: conn, X: 0, Y: 0, Name: conn.RemoteAddr().String(), Score: 0}
	playersMutex.Lock()
	players[conn] = player
	playersMutex.Unlock()
	broadcast(fmt.Sprintf("join#%s#%d#%d#%d", player.Name, player.X, player.Y, 0))

	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}
		
		command := strings.TrimSpace(string(data[:n]))
	 	if strings.HasPrefix(command, "eat#") {
			command = strings.Split(command, "#")[1]
			broadcast(command)
			placeFood()
			broadcast(fmt.Sprintf("food#%d#%d", foodX, foodY))
		} else {
			broadcast(command)
		}
	}

	playersMutex.Lock()	
	delete(players, conn)
	playersMutex.Unlock()
	broadcast(fmt.Sprintf("%s left the game.", player.Name))
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