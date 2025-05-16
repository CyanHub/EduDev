package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 15
)

type Game struct {
	conn         net.Conn
	x, y         int
	snake        [][2]float64
	foodX, foodY int
	mu           sync.Mutex
}

type Player struct {
	X, Y  int
	Name  string
	Score int
	Snake [][2]float64
}

var players = make(map[string]*Player)
var mu sync.Mutex

var client = &Player{}

func (g *Game) Update() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	// 记录移动前的位置
	oldX, oldY := g.x, g.y

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.y -= gridSize
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.y += gridSize
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		g.x -= gridSize
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		g.x += gridSize
	}

	// 只在实际移动时更新蛇的位置
	if oldX != g.x || oldY != g.y {
		// 更新蛇的位置
		if len(g.snake) > 0 {
			// 从尾部向前移动
			newSnake := make([][2]float64, len(g.snake))
			newSnake[0] = [2]float64{float64(g.x), float64(g.y)}
			copy(newSnake[1:], g.snake[:len(g.snake)-1])
			g.snake = newSnake
			client.X = g.x
			client.Y = g.y
			client.Score = len(g.snake)
			client.Snake = g.snake
			message, _ := json.Marshal(client)
			g.sendMessage(string(message))
		}
	}

	// 检查是否吃到食物
	if g.x == g.foodX && g.y == g.foodY {
		g.growSnake()
        // 消息延迟会导致一直吃到食物
        // g.foodX = -1
		// g.foodY = -1
		message, _ := json.Marshal(client)
		g.sendMessage("eat#" + string(message))
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.mu.Lock()
	mu.Lock()
	defer g.mu.Unlock()
	defer mu.Unlock()
	// 绘制蛇
	for i, pos := range g.snake {
		snakeImg := ebiten.NewImage(gridSize, gridSize)
		// 头部用不同颜色
		if i == 0 {
			vector.DrawFilledRect(snakeImg, 0, 0, float32(gridSize), float32(gridSize), color.RGBA{0, 200, 0, 255}, true)
		} else {
			vector.DrawFilledRect(snakeImg, 0, 0, float32(gridSize), float32(gridSize), color.RGBA{0, 255, 0, 255}, true)
		}

		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(pos[0], pos[1])
		screen.DrawImage(snakeImg, opts)
	}

	for _, player := range players {
		for _, pos := range player.Snake {
			snakeImg := ebiten.NewImage(gridSize, gridSize)
			vector.DrawFilledRect(snakeImg, 0, 0, float32(gridSize), float32(gridSize), color.RGBA{0, 255, 255, 255}, true)
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(pos[0], pos[1])
			screen.DrawImage(snakeImg, opts)
		}
	}

	// 绘制食物
	foodImg := ebiten.NewImage(gridSize, gridSize)
	vector.DrawFilledRect(foodImg, 0, 0, float32(gridSize), float32(gridSize), color.RGBA{255, 0, 0, 255}, true)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(g.foodX), float64(g.foodY))
	screen.DrawImage(foodImg, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) growSnake() {
	// 在尾部添加新段
	lastPos := g.snake[len(g.snake)-1]
	g.snake = append(g.snake, lastPos)
	client.Snake = g.snake
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		conn:  conn,
		x:     0,
		y:     0,
		snake: [][2]float64{{0, 0}},
	}
	client.Name = conn.LocalAddr().String()
	client.Snake = [][2]float64{{0, 0}}
	client.X = 0
	client.Y = 0
	client.Score = 0
	message, _ := json.Marshal(client)
	game.sendMessage(string(message))
	go game.handleMessage(conn)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) sendMessage(message string) {
	_, err := fmt.Fprintln(g.conn, message)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) handleMessage(conn net.Conn) {
	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		message := string(data[:n])
		message = strings.TrimSpace(message)
		mu.Lock()
		command := strings.Split(message, "#")
		if command[0] == "food" {
			fmt.Println("收到食物信息：", message)
			foodX, _ := strconv.Atoi(command[1])
			foodY, _ := strconv.Atoi(command[2])
			g.foodX = foodX
			g.foodY = foodY
		} else {
			var player Player
			json.Unmarshal([]byte(message), &player)
			if client.Name != player.Name {
				players[player.Name] = &player
			}
		}
		mu.Unlock()
	}
}
