package main

import (
	"log"
	"net"
	"io"
	"time"
	"strings"
)

// 客户端，表示连接
type Client struct {
	Name 	  string		// 名称
	Addr	  string		// 地址
	InfoChan  chan string	// 专属的信息通道
}

var (
	onlineMap map[string]Client	// 全局在线连接列表
	message = make(chan string)	// 全局的消息通道
)

// 固定格式构造消息
func MakeMsg(client Client, info string) string {
	buf := "[" + client.Addr + "]" + client.Name + ":" + info
	return buf
}

// 写信息到指定连接
func WriteToClient(client Client, conn net.Conn) {
	for msg := range client.InfoChan {
		conn.Write([]byte(msg + "\n"))
	}
}

// 处理连接
func HandlerConnect(conn net.Conn) {
	defer conn.Close()	// 延迟关闭
	
	quitChan := make(chan bool)	// 退出标记通道
	hasData := make(chan bool)	// 用户活跃标记通道

	addr := conn.RemoteAddr().String()	// 连接地址
	client := Client{
		Name: 		addr,
		Addr: 		addr,
		InfoChan: 	make(chan string),
	}
	onlineMap[addr] = client	// 添加到在线连接列表

	// 开启单独的go程写数据到client
	go WriteToClient()
	
	// 写入新连接login的消息到全局消息通道中
	message <- MakeMsg(client, "login")

	// 该go程用于保持和客户端的通信
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				quitChan <- true	// 读出0表示client退出，入队一个标记到退出通道中
				fmt.Println("Client Exit ...")
				return
			}
			
			if err != nil && err != io.EOF {	// 正常错误
				log.Printf("conn.Read Error: %v\n", err)
				return
			}

			// 与client对话
			msg := string(buf[:n - 1])
			if msg == "who" && len(msg) == 3 {	// 获取在线列表
				conn.Write([]byte("online user list: "))
				for _, v := onlineMap {
					conn.Write([]byte(v.Addr + ":" + v.Name + "\n"))
				}
			} else if len(strings.Split(msg, "|")) == 2 {	// 改名
				name := strings.Split(msg, "|")[1]
				client.Name = name
				onlineMap[addr] = client
				conn.Write([]byte("update name success"))
			} else {
				message <- MakeMsg(client, msg)	// 其他消息
			}

			hasData <- true // 若对话顺利进行则代表用户活跃，入队标记重置select的定时器
		}
	}()

	// select多路复用
	for {
		select {
			case <-quitChan:	// 退出标记
				delete(onlineMap, client.Addr)
				message <- MakeMsg(client, "logout")
				return
			case <-hasData:
				// 若用户活跃，重置此次select的定时器
			case <-time.After(time.Second * 10):	// 该连接超时
				delete(onlineMap, client.Addr)
				message <- MakeMsg(client, "logout")
				return
		}
	}
}

// 管理者
func Manager() {
	// 在线的连接列表
	onlineMap = make(map[string]Client)
	
	for {
		// 每当从全局消息通道中读出消息，就变量在线连接列表，将消息发送到每一个连接的专属通道中去
		msg := <- message
		for _, v := range onlineMap {
			v.InfoChan <- msg
		}
	}
}

func main() {
	// 端口监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Printf("net.Listen Error: %v\n", err)
		return
	}
	defer listener.Close()

	// 负责管理的go程
	go Manager()

	for {
		// 接收连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listener.Accept Error: %v\n", err)
			return
		}

		// 每收到一个连接就开启一个专属的go程处理
		go HandlerConnect(conn)
	}
}