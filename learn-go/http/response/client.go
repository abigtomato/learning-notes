package main

import (
	"fmt"
	"net"
)

/*
	模拟浏览器的http客户端
 */

/*
	HTTP响应报头格式:
	1.响应行: 协议/版本 + 空格 + 状态码 + 空格 + 状态码描述 \r\n
	2.响应头: Key: Value 一个或多个键值对 + \r\n
	3.空行: \r\n
	4.响应体: 请求的内容(返回HTML页面)

	成功实例:
	HTTP/1.1 200 OK
	Date: Tue, 13 Nov 2018 13:47:47 GMT
	Content-Length: 12
	Content-Type: text/plain; charset=utf-8

	Hello Golang

	失败实例:
	HTTP/1.1 404 Not Found
	Content-Type: text/plain; charset=utf-8
	X-Content-Type-Options: nosniff
	Date: Tue, 13 Nov 2018 13:49:27 GMT
	Content-Length: 19

	404 page not found
 */

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("net.Dial Error: %v\n", err)
		return
	}
	defer conn.Close()

	req_method := "GET /itcast HTTP/1.1\r\n"
	req_head := "Host:127.0.0.1:8000\r\n"
	req_space := "\r\n"
	
	_, err = conn.Write([]byte(req_method + req_head + req_space))
	if err != nil {
		fmt.Printf("conn.Write Error: %v\n", err)
		return
	}

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("conn.Read Error: %v\n", err)
		return
	}
	fmt.Printf("%s\n", buf[:n])
}