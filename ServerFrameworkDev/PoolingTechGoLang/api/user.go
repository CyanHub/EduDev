package api

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"

	"ServerFramework/global"
	"ServerFramework/model"
	"ServerFramework/model/request"
	"ServerFramework/model/response"
	"ServerFramework/service"
	"ServerFramework/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func calculateSum(n int) int {
	sum := 0
	for i := 0; i < n; i++ {  // 模拟业务输出逻辑
		sum += i
	}
	return sum
}

func Login(c *gin.Context) {

	for i := 0; i < 10; i++ {
		calculateSum(1000000000)
	}

	var req request.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Error("参数错误：", zap.String("err", utils.Translate(err)))
		response.FailWithMessage(utils.Translate(err), c)
		return
	}
	user, err := service.UserServiceApp.Login(req)
	if err != nil {
		if errors.Is(err, global.ErrUserNotFound) || errors.Is(err, global.ErrPasswordIncorrect) {
			global.Logger.Error("登陆失败：", zap.String("err", err.Error()))
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			log.Println("登录失败: ", err)
			response.FailWithMessage("登录失败", c)
			return
		}
	}
	jwt := utils.NewJwt()
	claims := jwt.CreateClaims(model.BaseClaims{
		UserId:   user.ID,
		Username: user.Username,
	})
	token, err := jwt.GenerateToken(&claims)
	if err != nil {
		log.Println("生成token失败: ", err)
		response.FailWithMessage("生成token失败", c)
		return
	}
	global.Logger.Info("登陆成功", zap.String("username", user.Username))
	response.OkWithData(&response.LoginResponse{
		User:  user,
		Token: token,
	}, c)
}

func Register(c *gin.Context) {
	var req request.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("参数错误: ", utils.Translate(err))
		response.FailWithMessage(utils.Translate(err), c)
		return
	}
	user, err := service.UserServiceApp.Register(req)
	if err != nil {
		if errors.Is(err, global.ErrUserAlreadyExists) {
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			log.Println("注册失败: ", err)
			response.FailWithMessage("注册失败", c)
			return
		}
	}
	response.OkWithData(user, c)
}

func UserList(c *gin.Context) {
	var req request.UserListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("参数错误: ", utils.Translate(err))
		response.FailWithMessage("参数错误: ", c)
		return
	}
	total, users, err := service.UserServiceApp.UserList(req)
	if err != nil {
		log.Println("获取用户列表失败: ", err)
		response.FailWithMessage("获取用户列表失败", c)
		return
	}
	response.OkWithData(response.PageResult{
		Total:    total,
		List:     users,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, c)
}

var clients = make(map[*websocket.Conn]string)
var broadcast = make(chan Message)
var mu = sync.Mutex{}

var upgrader = websocket.Upgrader{
	HandshakeTimeout: 10 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func OnlineTool(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Logger.Error("升级为WebSocket连接失败:", zap.Error(err))
		return
	}

	mu.Lock()
	clients[conn] = conn.RemoteAddr().String()
	mu.Unlock()

	go HandleClient(conn)
}

func HandleClient(conn *websocket.Conn) {
	defer func() {
		conn.Close()
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				global.Logger.Error("读取消息错误:", zap.Error(err))
			}
			return
		}

		message := string(p)
		global.Logger.Info("收到消息:", zap.String("from", clients[conn]), zap.String("message", message))

		Broadcast(messageType, p, conn)
	}
}

func Broadcast(messageType int, message []byte, sender *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	for conn := range clients {
		if conn != sender {
			err := conn.WriteMessage(messageType, message)
			if err != nil {
				global.Logger.Error("发送消息错误:", zap.Error(err))
				conn.Close()
				delete(clients, conn)
			}
		}
	}

}

func HandleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Logger.Error("Websocket升级失败:", zap.Error(err))
		return
	}
	defer ws.Close()

	mu.Lock()
	clients[ws] = ws.RemoteAddr().String()
	mu.Unlock()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}
		broadcast <- msg // 现在可以直接发送 Message 类型
	}
}

func BroadcastMessages() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				global.Logger.Error("发送消息错误:", zap.Error(err))
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

