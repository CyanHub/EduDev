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


func Login(c *gin.Context) {
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

// package api

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"sync"
// 	"time"

// 	"go.uber.org/zap"

// 	"ServerFramework/global"
// 	"ServerFramework/model"
// 	"ServerFramework/model/request"
// 	"ServerFramework/model/response"
// 	"ServerFramework/service"
// 	"ServerFramework/utils"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/websocket"
// )

// func Login(c *gin.Context) {
// 	var req request.UserLoginRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		//log.Println("参数错误: ", utils.Translate(err))
// 		global.Logger.Error("参数错误：", zap.String("err", utils.Translate(err))) // 新增
// 		response.FailWithMessage(utils.Translate(err), c)
// 		return
// 	}
// 	user, err := service.UserServiceApp.Login(req)
// 	if err != nil {
// 		if errors.Is(err, global.ErrUserNotFound) || errors.Is(err, global.ErrPasswordIncorrect) {
// 			global.Logger.Error("登陆失败：", zap.String("err", err.Error())) // 新增
// 			response.FailWithMessage(err.Error(), c)
// 			return
// 		} else {
// 			log.Println("登录失败: ", err)
// 			response.FailWithMessage("登录失败", c)
// 			return
// 		}
// 	}
// 	// 生成token
// 	jwt := utils.NewJwt()
// 	claims := jwt.CreateClaims(model.BaseClaims{
// 		UserId:   user.ID,
// 		Username: user.Username,
// 	})
// 	token, err := jwt.GenerateToken(&claims)
// 	if err != nil {
// 		log.Println("生成token失败: ", err)
// 		response.FailWithMessage("生成token失败", c)
// 		return
// 	}
// 	// 将用户信息缓存到`redis`中，对应的操作应该是`HASH`。
// 	//userJSON, err := json.Marshal(user)
// 	//if err != nil {
// 	//	log.Println("序列化用户信息失败: ", err)
// 	//	return
// 	//}
// 	//err = global.Redis.HSet(context.Background(), "online_user", user.ID, userJSON).Err()
// 	//if err != nil {
// 	//	log.Println("缓存用户信息失败: ", err)
// 	//}
// 	global.Logger.Info("登陆成功", zap.String("username", user.Username)) // 新增
// 	response.OkWithData(&response.LoginResponse{
// 		User:  user,
// 		Token: token,
// 	}, c)
// }

// func Register(c *gin.Context) {
// 	var req request.UserRegisterRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		log.Println("参数错误: ", utils.Translate(err))
// 		response.FailWithMessage(utils.Translate(err), c)
// 		return
// 	}
// 	user, err := service.UserServiceApp.Register(req)
// 	if err != nil {
// 		if errors.Is(err, global.ErrUserAlreadyExists) {
// 			response.FailWithMessage(err.Error(), c)
// 			return
// 		} else {
// 			log.Println("注册失败: ", err)
// 			response.FailWithMessage("注册失败", c)
// 			return
// 		}
// 	}
// 	response.OkWithData(user, c)
// }

// func UserList(c *gin.Context) {
// 	var req request.UserListRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		log.Println("参数错误: ", utils.Translate(err))
// 		response.FailWithMessage("参数错误: ", c)
// 		return
// 	}
// 	total, users, err := service.UserServiceApp.UserList(req)
// 	if err != nil {
// 		log.Println("获取用户列表失败: ", err)
// 		response.FailWithMessage("获取用户列表失败", c)
// 		return
// 	}
// 	response.OkWithData(response.PageResult{
// 		Total:    total,
// 		List:     users,
// 		Page:     req.Page,
// 		PageSize: req.PageSize,
// 	}, c)
// }

// // // 在线工具
// // type Message struct {
// // 	Type    int
// // 	Content string
// // }

// var broadcast = make(chan []byte)

// func HandleWebSocket(c *gin.Context) {
//     ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//     if err != nil {
//         global.Logger.Error("Websocket升级失败:", zap.Error(err))
//         return
//     }
//     defer ws.Close()

//     // clients[ws] = true

//     for {
//         _, message, err := ws.ReadMessage()
//         if err != nil {
//             delete(clients, ws)
//             break
//         }
//         broadcast <- message
//     }
// }

// func BroadcastMessages() {
//     for {
//         message := <-broadcast
//         for client := range clients {
//             err := client.WriteMessage(websocket.TextMessage, message)
//             if err != nil {
//                 global.Logger.Error("错误广播消息:", zap.Error(err))
//                 client.Close()
//                 delete(clients, client)
//             }
//         }
//     }
// }

// //////////////////////////////////////////////////////////////////
// var clients = make(map[*websocket.Conn]string)
// var mu = sync.Mutex{} // 互斥锁

// var upgrader = websocket.Upgrader{
// 	HandshakeTimeout: 10 * time.Second,
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func OnlineTool(c *gin.Context) {
// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		fmt.Println("升级为WebSocket连接失败,原因是:", err.Error())
// 	}

// 	mu.Lock()
// 	clients[conn] = conn.RemoteAddr().String() // 存储客户端连接
// 	mu.Unlock()
// 	go HandleClient(conn)
// }

// func HandleClient(conn *websocket.Conn) {
// 	defer func() {
// 		conn.Close()
// 		mu.Lock()
// 		delete(clients, conn)
// 		mu.Unlock()
// 	}()
// 	for {
// 		// 1.读取客户端发送的消息
// 		_, data, err := conn.ReadMessage()
// 		if err != nil {
// 			fmt.Printf("读取%s的消息失败,原因是:%s\n", conn.RemoteAddr().String(), err.Error())
// 			return
// 		}
// 		// 2.将消息广播给所有客户端
// 		Broadcast(data, conn)
// 	}
// }

// func Broadcast(data []byte, conn *websocket.Conn) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	for conn := range clients {
// 		if conn != nil {
// 			err := conn.WriteMessage(websocket.TextMessage, data)
// 			if err != nil {
// 				fmt.Printf("广播消息给客户端%s失败,原因是:%s\n", conn.RemoteAddr().String(), err.Error())
// 				continue
// 			}
// 		}
// 	}
// }
