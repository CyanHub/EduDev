package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	fmt.Println("http server start,listen port 8080")
	http.ListenAndServe(":8080", nil)
}
