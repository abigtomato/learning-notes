package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Printf("net.Dial Error: %v\n", err)
		return
	}
	defer conn.Close()

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