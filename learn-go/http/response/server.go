package main

import (
	"net/http"
)

/*
	Http服务器
 */

func main() {
	// 
	http.HandleFunc("/itcast", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("Hello Golang"))
	})
	http.ListenAndServe("127.0.0.1:8000", nil)
}