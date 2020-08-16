package main

import (
	"fmt"
	"net"
)

func main() {
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Printf("net.ResolveUDPAddr Error: %v\n", err)
		return
	}
	fmt.Printf("")

	conn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Printf("net.ListenUDP Error: %v\n", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 4096)
		n, cltAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("conn.ReadFromUDP Error: %v\n", err)
			return
		}
		fmt.Printf("")

		go func() {
			dayTime := time.Now().String() + "\n"
			_, err := conn.WriteToUDP([]byte(daytime), cltAddr)
			if err != nil {
				fmt.Printf("conn.WriteToUDP Error: %v\n", err)
				return
			}
		}()
	}
}