package main

import (
	"fmt"
	"net"
	"io"
	"os"
)

func recvFile(conn net.Conn, fileName string) (err error) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("os.Create Error: %v\n", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			if err != nil && err == io.EOF {
				fmt.Println("客户端断开连接")
			} else {
				fmt.Printf("conn.Read Error: %v\n", err)
			}
			return
		}

		_, err = f.Write(buf[:n])
		if err != nil {
			fmt.Printf("f.Write Error: %v\n", err)
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("net.Listen Error: %v\n", err)
		return
	}
	defer listen.Close()

	conn, err := listen.Accept()
	if err != nil {
		fmt.Printf("listen.Accept Error: %v\n", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("conn.Read Error: %v\n", err)
		return
	}

	_, err := conn.Write([]byte("ok"))
	if err != nil {
		fmt.Printf("conn.Write Error: %v\n", err)
		return
	}

	if err := recvFile(conn, buf[:n]); err != nil {
		fmt.Printf("recvFile Error: %v\n", err)
		return
	}
}