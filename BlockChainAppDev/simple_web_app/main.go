package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	fmt.Println("http服务启动,监听端口：8080")
	http.ListenAndServe(":8080", nil)
}
