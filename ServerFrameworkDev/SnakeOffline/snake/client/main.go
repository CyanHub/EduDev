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
	"ServerFramework/snake/frame"
	"strconv"
	"strings"
	"sync"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 15
)

type Game struct {
	x, y         int
	snake        [][2]float64
	foodX, foodY int
	mu           sync.Mutex
	Conn         net.Conn
}

// 需要保存其他玩家的信息
type Player struct {
	X, Y  int
	Snake [][2]float64
	Name  string
}

var players = make(map[string]*Player)
var playerMu sync.Mutex
var mySelf = Player{}

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
			mySelf.Y = g.y
			mySelf.X = g.x
			mySelf.Snake = g.snake
		}
		data, _ := json.Marshal(&mySelf)
		//SendMessage(g.Conn, data)
		SendFrame(g.Conn, data, frame.TypePlayerMove)
	}

	//fmt.Println("蛇的位置", g.x, g.y)
	// 检查是否吃到食物
	if g.x == g.foodX && g.y == g.foodY {
		g.growSnake()
		data, _ := json.Marshal(&mySelf)
		message := fmt.Sprintf("eat#%s", string(data))
		//SendMessage(g.Conn, []byte(message))
		SendFrame(g.Conn, []byte(message), frame.TypeEat)
		g.foodY = -1
		g.foodX = -1
		//g.placeFood()
	}
	//time.Sleep(500 * time.Millisecond)
	return nil
}

//func SendMessage(conn net.Conn, message []byte) {
//
//	_, err := conn.Write(message)
//	if err != nil {
//		fmt.Println("发送消息失败：", err.Error())
//		return
//	}
//}

func SendFrame(conn net.Conn, message []byte, msgType uint8) {
	_, err := frame.WriteFrame(conn, msgType, message)
	if err != nil {
		fmt.Println("发送消息失败：", err.Error())
		return
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	playerMu.Lock()
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
	fmt.Println("增长")
	// 在尾部添加新段
	lastPos := g.snake[len(g.snake)-1]
	g.snake = append(g.snake, lastPos)
}

func main() {
	game := &Game{
		x:     0,
		y:     0,
		snake: [][2]float64{{0, 0}},
	}

	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	mySelf.X = game.x
	mySelf.Y = game.y
	mySelf.Name = conn.LocalAddr().String()
	mySelf.Snake = game.snake
	game.Conn = conn

	data, _ := json.Marshal(mySelf)
	//SendMessage(conn, data)
	SendFrame(conn, data, frame.TypePlayerMove)

	go game.HandlerMessage(conn)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) HandlerMessage(conn net.Conn) {
	// 1. 要接受客户端消息（其他客户端的位置）
	for {
		//var buffer = make([]byte, 2048)
		//n, err := conn.Read(buffer)
		messageFrame, err := frame.ReadFrame(conn)
		if err != nil {
			fmt.Println("读取消息失败：", err.Error())
		}
		message := string(messageFrame.Data)
		//message = strings.TrimSpace(message)
		//messageSlice := strings.Split(message, "#")
		switch messageFrame.Type {
		case frame.TypePlayerMove:
			player := Player{}
			err = json.Unmarshal([]byte(message), &player)
			if err != nil {
				fmt.Println("反序列化失败：", err)
				continue
			}
			if mySelf.Name != player.Name {
				fmt.Println("收到玩家信息：", player)
				playerMu.Lock()
				players[player.Name] = &player
				playerMu.Unlock()
			}
			break
		case frame.TypeFood:
			fmt.Println("收到食物消息：", message)
			messageSlice := strings.Split(message, "#")
			foodX, _ := strconv.Atoi(messageSlice[1])
			foodY, _ := strconv.Atoi(messageSlice[2])
			g.foodX = foodX
			g.foodY = foodY
			fmt.Println("食物的位置", g.foodX, g.foodY)
			break
		}
	}
}
