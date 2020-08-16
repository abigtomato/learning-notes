package main

import (
	"fmt"
	"net"
	"io"
	"os"
)

func sendFile(conn net.Conn, filePath string) error {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%v\n", err)
			} else {
				fmt.Printf("f.Read Error: %v\n", err)
			}
			return
		}

		_, err := conn.Write(buf[:n])
		if err != nil {
			fmt.Printf()
			return
		}
	}
}	

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("格式错误")
		return 
	}

	filePath := list[1]
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("os.Stat Error: %v\n", err)
		return
	}
	fileName := fileInfo.Name()

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("net.Dial Error: %v\n", err)
		return
	}
	defer conn.Close()

	_, err := conn.Write([]byte(fileName))
	if err != nil {
		fmt.Printf("conn.Write Error: %v\n", err)
		return
	}

	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("conn.Read Error: %v\n", err)
		return
	}

	if "ok" == string(buf[:n]) {
		sendFile(filePath)
	} else {
		fmt.Printf("Ok Error: %v\n", err)
		return
	}
}