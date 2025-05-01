package main

import (
	"BlockChainDev/web_server/config"
	"BlockChainDev/web_server/pkg/logs"
	"BlockChainDev/web_server/pkg/mysqldb"
	"BlockChainDev/web_server/pkg/redisdb"
	"BlockChainDev/web_server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置文件
	config.Init()
	// 初始化数据库链接（MySQL）
	mysqldb.InitMysql()
	// 初始化日志
	logs.InitLogger(
		config.CONFIG.Logger.LogTypes,
		config.CONFIG.Logger.Dir,
		logs.LogEnvType(config.CONFIG.System.Mode),
		config.CONFIG.Logger.LogMaxAge,
	)
	// 初始化Redis链接
	redisdb.Init()

	//// 启动协程 生成红包
	//go func() {
	//	produce_redpack.Init() // 调用 produce_redpack 包的 Init 函数
	//}()

	gin.SetMode(config.CONFIG.System.Mode)
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// 初始化数据库表路由
	router.InitUserRouter(r)
	err := r.Run(":" + config.CONFIG.System.Port)
	if err != nil {
		panic(err)
	}
	// 初始化红包相关路由
	//router.InitRedpackRouter(r)
	//err := r.Run(":" + config.CONFIG.System.Port)
	//if err != nil {
	//	logs.ZapLogger.Error("启动服务器失败: " + err.Error())
	//	panic(err)
	//}
}

//package main
//
//import (
//	"BlockChainDev/redis_server/config"
//	//"BlockChainDev/redis_server/internal/produce_redpack"
//	"BlockChainDev/redis_server/pkg/logs"
//	"BlockChainDev/redis_server/pkg/mysqldb"
//	"BlockChainDev/redis_server/pkg/redisdb"
//	"BlockChainDev/redis_server/router"
//	"github.com/gin-gonic/gin"
//)
//
//func main() {
//	// 初始化配置文件
//	config.Init()
//	// 初始化数据库链接（MySQL）
//	mysqldb.InitMysql()
//	// 初始化日志
//	logs.InitLogger(
//		config.CONFIG.Logger.LogTypes,
//		config.CONFIG.Logger.Dir,
//		logs.LogEnvType(config.CONFIG.System.Mode),
//		config.CONFIG.Logger.LogMaxAge,
//	)
//	// 初始化Redis链接
//	redisdb.Init()
//
//	gin.SetMode(config.CONFIG.System.Mode)
//	r := gin.Default()
//	r.SetTrustedProxies(nil)
//	// 初始化红包相关路由
//	router.InitRedpackRouter(r)
//	err := r.Run(":" + config.CONFIG.System.Port)
//	if err != nil {
//		logs.ZapLogger.Error("启动服务器失败: " + err.Error())
//		panic(err)
//	}
//
//	////创建红包接口使用的是 POST 请求，而浏览器地址栏只能发送 GET 请求，所以不能直接在浏览器地址栏测试该接口。编写一个简单的程序来发送 POST 请求。
//	//url := "http://127.0.0.1:8188/api/v1/redpack/"
//	//// 要发送的 JSON 数据
//	//data := []byte(`{"amount": 100, "num": 10}`)
//	//
//	//// 创建一个新的 HTTP 请求
//	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
//	//if err != nil {
//	//	fmt.Println("创建请求失败:", err)
//	//	return
//	//}
//	//
//	//// 设置请求头
//	//req.Header.Set("Content-Type", "application/json")
//	//
//	//// 发送请求
//	//client := &http.Client{}
//	//resp, err := client.Do(req)
//	//if err != nil {
//	//	fmt.Println("发送请求失败:", err)
//	//	return
//	//}
//	//defer resp.Body.Close()
//	//
//	//// 读取响应内容
//	//body, err := ioutil.ReadAll(resp.Body)
//	//if err != nil {
//	//	fmt.Println("读取响应内容失败:", err)
//	//	return
//	//}
//	//
//	//fmt.Println("响应状态码:", resp.StatusCode)
//	//fmt.Println("响应内容:", string(body))
//}
