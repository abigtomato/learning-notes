package main

import (
	"fmt"
	"net"
	"os"
	"log"
)

/*
	Socket实现Http服务器
 */

/*
	HTTP请求包:
	1.请求行: 请求方法 + 空格 + 请求URL + 空格 + 协议版本 + \r\n
	2.请求头: Key: Value 多个键值对的格式封装数据
	3.空行: \r\n
	4.请求包体: 请求方法对应的数据(POST)
 */

func errFunc(err error, errInfo string) {
	if err != nil {
		log.Printf("%s Error: %v\n", errInfo, err)
		os.Exit(1)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	errFunc(err, "net.Listen")
	defer listen.Close()

	conn, err := listen.Accept()
	errFunc(err, "listen.Accept")
	defer conn.Close()
	
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		errFunc(err, "conn.Read")
		fmt.Printf("Request Package: \r\n%s\n", buf[:n])
	}
}