package main

import (
	"fmt"
	"time"
)

/* 只读只写channel的使用 */

func main() {
	var intChan chan int
	var exitChan chan struct{}
	
	intChan = make(chan int, 10)
	exitChan = make(chan struct{}, 2)

	// chan<- 指定管道的状态为只写，适用于发送数据之类的场景
	go func(intChan chan<- int, exitChan chan struct{}) {
		defer close(intChan)
		
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			intChan <- i
			fmt.Printf("发送数据: %v\n", i)
		}

		var flag struct{}
		exitChan <- flag
	}(intChan, exitChan)

	// <-chan 指定管道的状态为只读，适用于接收数据之类的场景
	go func(intChan <-chan int, exitChan chan struct{}) {
		for {
			if val, ok := <-intChan; !ok {
				break				
			} else {
				fmt.Printf("接收数据: %v\n", val)
			}
		}

		var flag struct{}
		exitChan <- flag
	}(intChan, exitChan)

	/*
		1.主线程循环判断exitChan管道的结束标记；
		2.若存在2个结束标记则代表接收和发送协程都结束任务，之后主线程结束。
	*/
	var total = 0
	for {
		if _, ok := <-exitChan; ok {
			total++
		}	

		if total == 2 {
			break
		}
	}
}