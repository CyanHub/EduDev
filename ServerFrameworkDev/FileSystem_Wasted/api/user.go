package api

import (
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"

	"FileSystem/global"
	"FileSystem/model"
	"FileSystem/model/request"
	"FileSystem/model/response"
	"FileSystem/service"
	"FileSystem/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
)

func calculateSum(n int) int {
	sum := 0
	for i := 0; i < n; i++ { // 模拟业务输出逻辑
		sum += i
	}
	return sum
}

// Login 用户登录
func Login(c *gin.Context) {
	var req request.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}

	// 输入校验
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		response.FailWithMessage(utils.Translate(err), c)
		return
	}

	// 验证用户信息
	user, err := service.UserServiceApp.Login(req) // 修改为传入整个请求对象
	if err != nil {
		global.Logger.Error("登录失败", zap.Error(err))
		response.FailWithMessage("用户名或密码错误", c)
		return
	}

	// 生成JWT令牌
	claims := model.BaseClaims{
		UserID:   user.ID,
		Username: user.Username,
		RoleID:   user.RoleId,
	}
	token, err := utils.NewJWT().CreateToken(claims)
	if err != nil {
		global.Logger.Error("生成Token失败", zap.Error(err))
		response.FailWithMessage("登录失败", c)
		return
	}

	// 记录操作日志
	global.Logger.Info("用户登录",
		zap.String("username", user.Username),
		zap.Uint64("userId", user.ID))

		response.OkWithDetailed(gin.H{
			"token": token,
			"expire": time.Now().Add(time.Duration(global.CONFIG.Jwt.ExpireTime) * time.Second).Unix(),
			"user": gin.H{
				"username": user.Username,
				"role": user.RoleId,
			},
		}, "登录成功", c)
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

// func HandleWebSocket(c *gin.Context) {
// 	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		global.Logger.Error("Websocket升级失败:", zap.Error(err))
// 		return
// 	}
// 	defer ws.Close()

// 	mu.Lock()
// 	clients[ws] = ws.RemoteAddr().String()
// 	mu.Unlock()

// 	for {
// 		var msg Message
// 		err := ws.ReadJSON(&msg)
// 		if err != nil {
// 			mu.Lock()
// 			delete(clients, ws)
// 			mu.Unlock()
// 			break
// 		}
// 		broadcast <- msg // 现在可以直接发送 Message 类型
// 	}
// }

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

func Logout(c *gin.Context) {
	// 清除JWT token
	c.SetCookie("token", "", -1, "/", "", false, true)
	response.OkWithMessage("登出成功", c)
}
