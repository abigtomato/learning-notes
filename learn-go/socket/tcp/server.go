package main

import (
	"fmt"
	"net"
)

/* 
	TCP编程-服务器
 */

/*
	TCP通信过程:
	1.三次握手:
		1.1 连接发起方发送SYN标志位(携带一串数字如:1000，数据包大小如:0等信息)，
			此时主动端TCP处于SYN_SEND状态，被动端处于LISTEN状态；
		1.2 被动接收方接收后发送ACK应答包(将接收到的一串数字加1如:1001发送出去，表示1000接收完毕)和SYN标志位(携带一串数字如:8000，数据包大小如:0等信息)；
			此时被动端TCP处于SYN_RCVD状态；
		1.3 连接发起方收到应答后也向对方发送ACK应答包(将接收到的一串数字加1如:8000发送出去)，
			标志着TCP连接建立完成，对应Server端的Accept()函数返回，Client端的Dial()函数返回，
			此时主被动端的TCP都处于ESTABLISHED消息传输状态。
	2.通信过程:
		1.1 主动发送端发送ACK包，携带一串数字如:1001，数据包大小如:20和数据；
		1.2 被动接收端接收到数据包也会发送ACK应答包，将对方发送的一串数字和数据包大小相加为1021携带出去，
			同时一并携带一串数字如:8001，数据包大小如:10和数据；
		1.3 主动发送端接收到数据会再次发送ACK应答包，并携带对方发送的一串数字和数据包大小相加8011。
	3.四次挥手:
		1.1 主动关闭方发送FIN标志位，携带包头如:1021，数据包大小如:0，如果之前有ACK通信还需要携带ACK应答包，
			此时主动端TCP处于FIN_WAIT_1状态；
		1.2 被动关闭方接数据包后，发送ACK应答，并携带对方发送的包头加1(1022)，
			此时整个连接处于半关闭状态，主动关闭方已经不需要主动发送数据了，
			此时主动端TCP处于FIN_WAIT_2状态(半关闭状态)，被动端TCP处于CLOSE_WAIT状态；
		1.3 被动关闭方法发送FIN包，携带包头如:8011，数据包大小如:0，ACK应答和包头1022，
			此时主动端TCP处于TIME_WAIT状态，经过2MSL时长变回CLOSED状态(如果被动端未接收到FIN会再次发送，主动端等待时长就是为此而存在)，
			被动端TCP处于LAST_ACK状态；
		1.4 主动关闭方法只需要发送ACK应答即可(携带8012，即对方发送的包头加1)，
			被动端接收后处于CLOSED状态。
 */

func main() {
	// net.Listen(): 绑定服务端的ip地址和端口号
	// 参数1: 选用的协议，如: "tcp", "udp"
	// 参数2: ip地址 + 端口号
	// 返回值: 绑定ip端口用于监听的套接字
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Printf("net.Listen Error: %v\n", err)
		return
	}
	defer listen.Close()

	for {
		// listen.Accept(): 阻塞监听客户端的连接
		// 返回值: 与客户端通信的套接字
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("listen.Accept Error: %v\n", err)
		} else {
			fmt.Printf("与IP: %v 建立连接\n", conn.RemoteAddr().String())
		}

		// 为conn启动单独的go程服务
		go func() {
			// 等服务完毕关闭客户端连接
			defer conn.Close()

			for {
				buf := make([]byte, 4096)
				// Read(): 从套接字接收客户端发送的数据
				// 参数: 字节数组缓冲区
				// 返回值: 实际读取到的字节数和错误
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Printf("conn.Read Error: %v\n", err)
					return
				}
				// 对端管道关闭，则会读取出0
				if n == 0 {
					if err != nil && err == io.EOF {
						fmt.Println("客户端断开连接")
					} else {
						fmt.Printf("conn.Read Error: %v\n", err)
					}
					return
				}
				fmt.Printf("ip: %v -> [%v]\n", conn.RemoteAddr().String(), string(buf[:n]))
			}
		}()
	}
}