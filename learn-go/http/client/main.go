package main

import (
	"net/http"
	"fmt"
)

/* 
	type Response struct {
		Status 		string
		StatusCode  int
		Proto 		string
		......
		Header 		Header
		Body 		io.ReadCloser
		......
	}
 */

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Printf("http.Get Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Header: %v\n", resp.Header)
	fmt.Printf("Status: %v\n", resp.Status)
	fmt.Printf("StatusCode: %v\n", resp.StatusCode)
	fmt.Printf("Proto: %v\n", resp.Proto)

	buf := make([]byte, 4096)
	var result string 
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%v\n", result)
			} else {
				fmt.Printf("resp.Body.Read Error: %v\n", err)
			}
			return
		}
		result += string(buf[:n])
	}
}