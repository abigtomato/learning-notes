package main

import (
	"fmt"
	"time"
)

/* recover处理go程错误 */

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello, Golang")
	}
}

// 测试协程运行时出现错误的解决方案
func test() {
	// 使用defer和recover捕获goroutine出现的错误
	// 这样做的好处是避免协程运行中断导致整个程序运行结束
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err=%v\n", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "Golang"
}

func main() {
	go test()
	go sayHello()

	time.Sleep(time.Second * 15)
}