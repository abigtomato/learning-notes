package main

import (
	"fmt"
	"net"
	"os"
)

/* 
	tcp编程-客户端
 */

func main() {
	// net.Dial(): 与服务器建立连接
	// 参数: 协议和服务端的ip+端口
	// 返回值: 与服务器通信的套接字
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("与服务器建立连接失败，错误信息=%v\n", err)
		return 
	}
	defer conn.Close()

	// 单独go程不断接收服务器发送的数据
	go func() {
		buf := make([]byte, 4096)
		for {	
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Printf("conn.Read Error: %v\n", err)
				return
			}
			fmt.Printf("Server Request: %v\n", string(buf[:n]))
		}
	}()
	
	// 客户端主go程负责将标准输入流的数据发送给服务器端
	buf := make([]byte, 4096)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Printf("os.Stdin.Read Error: %v\n", err)
			continue
		}
		if string(buf[:n]) == "exit\r\n" {
			break
		}
		conn.Write(buf[:n])
	}
}