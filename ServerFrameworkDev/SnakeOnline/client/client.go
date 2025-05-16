package main

import (
	"ServerFrameWork/pkg"
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

type Player struct {
	X, Y  int
	Snake [][2]float64 // 蛇的位置
	Name  string       // 玩家名称
}

var players = make(map[string]*Player) // 保存所有玩家信息（连接）
var mu sync.Mutex                      // 保护players的读写操作
var mySelf = Player{}                  // 保存自己的信息

type Game struct {
	x, y         int
	snake        [][2]float64
	foodX, foodY int
	conn         net.Conn
	mu           sync.Mutex
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
			mySelf.Y = g.y
			mySelf.X = g.x
			mySelf.Snake = g.snake

			// 发送移动消息给服务器
			data, err := json.Marshal(&mySelf)
			if err != nil {
				fmt.Println("序列化失败", err.Error())
				return err
			}
			// g.SendMessage(data)
			g.SendFrame(data, pkg.TypePlayerMove)
		}

	}

	//fmt.Println("蛇的位置", g.x, g.y)
	// 检查是否吃到食物
	if g.x == g.foodX && g.y == g.foodY {
		g.growSnake()
		data, err := json.Marshal(&mySelf)
		if err != nil {
			fmt.Println("序列化失败", err.Error())
			return err
		}
		message := fmt.Sprintf("eat#%s", string(data))
		// g.SendMessage([]byte(message))
		g.SendFrame([]byte(message), pkg.TypeEat)
		// g.placeFood()
	}
	return nil
}

func (g *Game) SendFrame(data []byte, msgType uint8) {
	err := pkg.WriteFrame(g.conn, msgType, data)
	if err != nil {
		fmt.Println("发送消息失败", err.Error())
	}
}

// func (g *Game) SendMessage(data []byte) {
// 	_, err := g.conn.Write(data)
// 	if err != nil {
// 		fmt.Println("发送消息失败", err.Error())
// 	}
// }

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
}

func main() {
	game := &Game{
		x:     0,
		y:     0,
		snake: [][2]float64{{0, 0}},
	}

	conn, err := net.Dial("tcp", "localhost:5090") // 连接服务器
	if err != nil {
		panic(err)
	}
	game.conn = conn          // 保存连接
	mySelf.X = game.x         // 保存自己的X信息
	mySelf.Y = game.y         // 保存自己的Y信息
	mySelf.Snake = game.snake // 保存自己的蛇信息
	mySelf.Name = conn.LocalAddr().String()
	fmt.Println("连接服务器成功, 开始游戏...")

	go game.HandleServerMessage() // 处理服务器消息

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("贪吃蛇 联机版")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) HandleServerMessage() {
	defer func() { // 关闭连接
		g.conn.Close() // 关闭连接

	}()
	// 接收服务器发送的消息 放置事物的消息，其他玩家的位置的消息
	for {
		// buffer := make([]byte, 2048)  // 读取消息
		// n, err := g.conn.Read(buffer) // 读取消息

		messageFrame, err := pkg.ReadFrame(g.conn)
		if err != nil {
			fmt.Printf("客户端%s发送消息失败%s\n", g.conn.LocalAddr().String(), err.Error())
			return
		}
		// message := string(buffer[:n])               // 读取消息
		message := string(messageFrame.Data)        // 读取消息
		messageSlice := strings.Split(message, "#") // 分割消息
		switch messageFrame.Type {
		case pkg.TypePlayerMove:
			fmt.Println("收到来自服务器广播的玩家位置信息：", message)
			// 收到玩家位置信息
			player := Player{}                                      // 玩家信息
			err := json.Unmarshal([]byte(messageSlice[0]), &player) // 解析消息
			if err != nil {
				fmt.Println("反序列失败，解析消息失败：", err.Error())
				continue
			}
			if player.Name != mySelf.Name {
				mu.Lock()                      // 加锁，防止并发读写
				players[player.Name] = &player // 保存玩家连接
				mu.Unlock()                    // 解锁
			}
			// break  // 这个Break其实冗余了，有和没有是没区别的
		case pkg.TypeFood:
			// 收到服务器发来的食物信息
			x, _ := strconv.Atoi(messageSlice[1]) // 食物X坐标
			y, _ := strconv.Atoi(messageSlice[2]) // 食物Y坐标
			g.foodX = x                           // 保存食物X坐标
			g.foodY = y                           // 保存食物Y坐标
		}

		// if len(messageSlice) == 1 { // 分割消息
		// 	fmt.Println("收到来自服务器广播的玩家位置信息：", message)
		// 	// 收到玩家位置信息
		// 	player := Player{}                                      // 玩家信息
		// 	err := json.Unmarshal([]byte(messageSlice[0]), &player) // 解析消息
		// 	if err != nil {
		// 		fmt.Println("反序列失败，解析消息失败：", err.Error())
		// 		continue
		// 	}
		// 	if player.Name != mySelf.Name {
		// 		mu.Lock()                      // 加锁，防止并发读写
		// 		players[player.Name] = &player // 保存玩家连接
		// 		mu.Unlock()                    // 解锁
		// 	}
		// } else if len(messageSlice) == 3 { // 分割消息}
		// 	// 收到服务器发来的食物信息
		// 	x, _ := strconv.Atoi(messageSlice[1]) // 食物X坐标
		// 	y, _ := strconv.Atoi(messageSlice[2]) // 食物Y坐标
		// 	g.foodX = x                           // 保存食物X坐标
		// 	g.foodY = y                           // 保存食物Y坐标
		// }
		// }
	}
}
