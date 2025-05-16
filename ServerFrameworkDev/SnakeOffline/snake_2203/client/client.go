package main

import (
	"encoding/json"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"net"
	"ServerFramework/snake_2203/pkg"
	"strconv"
	"strings"
	"sync"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 15
)

type Player struct {
	X, Y  int
	Snake [][2]float64
	Name  string
}

var players = make(map[string]*Player)
var playerMu sync.Mutex
var mySelf = &Player{}

type Game struct {
	x, y         int
	snake        [][2]float64
	foodX, foodY int
	mu           sync.Mutex
	conn         net.Conn
}

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
			mySelf.X = g.x
			mySelf.Y = g.y
			mySelf.Snake = g.snake

			// 把位置信息发送给服务器
			data, err := json.Marshal(&mySelf)
			if err != nil {
				fmt.Println("序列化失败，原因是", err.Error())
			} else {
				g.SendFrame(data, pkg.TypePlayerMove)
			}
		}

	}

	//fmt.Println("蛇的位置", g.x, g.y)
	// 检查是否吃到食物
	if g.x == g.foodX && g.y == g.foodY {
		//fmt.Printf("g.x = %d, g.y = %d\n g.foodX = %d, g.foodY =  %d", g.x, g.y, g.foodX, g.foodY)
		g.growSnake()
		// 把吃到的信息发送给服务器
		data, err := json.Marshal(&mySelf)
		if err != nil {
			fmt.Println("序列化失败，原因是", err.Error())
		} else {
			message := "eat#" + string(data)
			g.SendFrame([]byte(message), pkg.TypeEat)
		}
		//g.placeFood()
	}
	return nil
}

func (g *Game) SendFrame(message []byte, msgType uint8) {
	err := pkg.WriteFrame(g.conn, msgType, message)
	if err != nil {
		fmt.Println("发送消息失败：", err.Error())
		return
	}
}

//
//func (g *Game) SendMessage(message []byte) {
//	_, err := g.conn.Write(message)
//	if err != nil {
//		fmt.Println("发送消息给服务器失败，原因是：", err.Error())
//	}
//}

func (g *Game) Draw(screen *ebiten.Image) {
	g.mu.Lock()
	playerMu.Lock()
	defer g.mu.Unlock()
	defer playerMu.Unlock()
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
}

//func (g *Game) placeFood() {
//	g.foodX = (rand.Intn(screenWidth / gridSize)) * gridSize
//	g.foodY = (rand.Intn(screenHeight / gridSize)) * gridSize
//	fmt.Println("Food placed")
//}

func main() {
	conn, err := net.Dial("tcp", ":10086")
	if err != nil {
		panic(err)
	}

	game := &Game{
		x:     0,
		y:     0,
		snake: [][2]float64{{0, 0}},
		conn:  conn,
		foodX: -1,
		foodY: -1,
	}

	mySelf.X = game.x
	mySelf.Y = game.y
	mySelf.Snake = game.snake
	mySelf.Name = conn.LocalAddr().String()

	// 1. 接受服务器传来的消息（食物位置、其他玩家位置）
	// 2. 每次移动给服务器发送新的位置
	// 3. 吃到食物的时候要告知服务器

	go game.handleServerMessage(conn)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) handleServerMessage(conn net.Conn) {

	defer conn.Close()
	for {

		messageFrame, err := pkg.ReadFrame(conn)

		//n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("接收服务器消息失败，原因是：", err.Error())
			return
		}
		message := string(messageFrame.Data)
		messageSlice := strings.Split(message, "#")
		switch messageFrame.Type {
		case pkg.TypePlayerMove:
			// 玩家状态消息
			player := Player{}
			err := json.Unmarshal(messageFrame.Data, &player)
			if err != nil {
				fmt.Println("反序列化失败，原因是：", err.Error())
				continue
			}
			if player.Name != mySelf.Name {
				playerMu.Lock()
				players[player.Name] = &player
				playerMu.Unlock()
			}
			break
		case pkg.TypeFood:
			fmt.Println("收到食物位置信息", messageSlice[1], messageSlice[2])
			// 服务器放置食物的信息
			x, _ := strconv.Atoi(messageSlice[1])
			y, _ := strconv.Atoi(messageSlice[2])
			g.foodX = x
			g.foodY = y
		}

	}
}
