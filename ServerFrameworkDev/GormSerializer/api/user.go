package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"go.uber.org/zap"
	"golang.org/x/net/websocket"

	"ServerFramework/global"
	"ServerFramework/model"
	"ServerFramework/model/request"
	"ServerFramework/model/response"
	"ServerFramework/service"
	"ServerFramework/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req request.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//log.Println("参数错误: ", utils.Translate(err))
		global.Logger.Error("参数错误：", zap.String("err", utils.Translate(err))) // 新增
		response.FailWithMessage(utils.Translate(err), c)
		return
	}
	user, err := service.UserServiceApp.Login(req)
	if err != nil {
		if errors.Is(err, global.ErrUserNotFound) || errors.Is(err, global.ErrPasswordIncorrect) {
			global.Logger.Error("登陆失败：", zap.String("err", err.Error())) // 新增
			response.FailWithMessage(err.Error(), c)
			return
		} else {
			log.Println("登录失败: ", err)
			response.FailWithMessage("登录失败", c)
			return
		}
	}
	// 生成token
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
	// 将用户信息缓存到`redis`中，对应的操作应该是`HASH`。
	//userJSON, err := json.Marshal(user)
	//if err != nil {
	//	log.Println("序列化用户信息失败: ", err)
	//	return
	//}
	//err = global.Redis.HSet(context.Background(), "online_user", user.ID, userJSON).Err()
	//if err != nil {
	//	log.Println("缓存用户信息失败: ", err)
	//}
	global.Logger.Info("登陆成功", zap.String("username", user.Username)) // 新增
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
		response.FailWithMessage(utils.Translate(err), c)
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

// // // // // //
// type Message struct {
// 	Type    string `json:"type"`
// 	Content string `json:"content"`
// }

var clients = make(map[*websocket.Conn]string)
var mu sync.Mutex

var upgrader = websocket.upgrader{
	HandShakeTimeout: 10 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func OnlineTool(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("建立WebSocket连接失败，失败的原因为：", err.Error())
		// return
	}
	mu.Lock()
	clients[conn] = conn.RemoteAddr().String()
	mu.Unlock()
	go HandleClienet(conn)

}

func HandleClienet(conn *websocket.Conn) {
	defer func() {
		conn.Close()
		mu.Lock()
		delete(clients, conn)
		mu.Unlock()
		// conn.Close()
	}()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("建立%s连接失败，失败的原因为：%s\n", conn.RemoteAddr().String(), err.Error())
			return
		}
		// 处理接收到的消息，广播出去(发送给所有客户端)
		Broadcast(data)
	}
}

func Broadcast(data []byte, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	for conn := range clients {
		conn.WriteMessage(websocket.TextMessage, data)
		err := nil
		if err != nil {
			fmt.Printf("向%s发送消息失败，失败的原因为：%s\n", conn.RemoteAddr().String(), err.Error())
			continue
		}
	}
}
