package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func OpenSendFile(url string, resp http.ResponseWriter) {
	path := "./file/" + url
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		resp.Write([]byte("No such file or directory !"))
		return 
	}

	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Read Out.")
			} else {
				fmt.Printf("f.Read Error: %v\n", err)
			}
			return
		}
		resp.Write(buf[:n])
	}
}

func main() {
	/*
		type ResponseWriter interface {
			Header() Header
			Write([]byte) (int, err)
			WriteHeader(int)
		}

		type Request struct {
			Method 	   string
			URL 	   *url.URL
			......
			Header 	   Header
			Body 	   io.ReadCloser
			RemoteAddr string
			......
			ctx 	   context.Context
		}
	 */
	// http.HandleFunc: 为一个请求的url注册处理器(回调函数)
	// 参数1: 请求的URL
	// 参数2: func(resp http.ResponseWriter, req *http.Request) 类型的回调函数
	http.HandleFunc("/itcast", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Printf("URL: %v\n", req.URL)
		fmt.Printf("Method: %v\n", req.Method)
		fmt.Printf("Host: %v\n", req.Host)
		fmt.Printf("RemoteAddr: %v\n", req.RemoteAddr)
		fmt.Printf("Body: %v\n", req.Body)

		OpenSendFile(req.URL, resp)		
	})

	// http.ListenAndServe: 绑定服务器端ip + port和注册回调函数
	// 参数1: 绑定服务端IP地址和端口
	// 参数2: 默认回调 http.DefaultServeMux() 处理
	http.ListenAndServe("127.0.0.1:8000", nil)
}