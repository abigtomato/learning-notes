package main

import (
	"fmt"
	"net"
	"log"
)

/* 聊天服务器示例 */

type client chan<- string

var (
	entering = make(chan client)	// 客户进入通道
	leaving = make(chan client)		// 客户离开通道
	messages = make(chan string)	// 消息通道
)

// 广播go程
func broadcaster() {
	// 保存所有连接的客户端
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-message:
			// 将收到的消息广播给所有客户端
			for cli, _ := range clients {
				cli <- msg
			}
		case cli := <-entering:
			// 新客户端加入
			clients[cli] = true
		case cli := <-leaving:
			// 客户端离开
			delete(clients, cli)
			close(cli)
		}
	}
}

// 处理连接go程
func handleConn(conn net.Conn) {
	defer conn.Close()
	
	// 每个客户端连接单独的通道
	ch := make(chan string)
	go clientWriter(conn, ch)
	
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived."	// 进行广播
	entering <- ch	// 通知新客户到来

	// 读取客户端发送的所有消息
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch	// 通知客户的离开
	messages <- who + " has left."	// 进行广播
}

// 写消息到客户端的go程
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	// 监听IP和端口
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// 开启广播go程
	go broadcaster()

	for {
		// 接收客户连接
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		
		// 开启连接处理go程
		go handleConn(conn)
	}
}