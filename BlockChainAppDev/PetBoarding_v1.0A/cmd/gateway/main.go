package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 用户服务路由
	user := createReverseProxy("http://localhost:8081")
	r.Any("/user/*path", gin.WrapH(user))

	// 宠物服务路由
	pet := createReverseProxy("http://localhost:8082")
	r.Any("/pet/*path", gin.WrapH(pet))

	// 订单服务路由
	order := createReverseProxy("http://localhost:8084")
	r.Any("/order/*path", gin.WrapH(order))

	// 评价服务路由
	review := createReverseProxy("http://localhost:8083")
	r.Any("/review/*path", gin.WrapH(review))

	fmt.Println("API网关服务启动中，监听端口:8080...")
	r.Run(":8080")
}

func createReverseProxy(target string) *httputil.ReverseProxy {
	targetUrl, _ := url.Parse(target)
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = targetUrl.Scheme
			req.URL.Host = targetUrl.Host
			req.URL.Path = targetUrl.Path + req.URL.Path
		},
	}
}
