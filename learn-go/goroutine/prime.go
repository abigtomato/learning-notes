package main

import (
	"fmt"
	"time"
	"runtime"
)

/* 生产消费模型计算素数 */

// 生产者
func ProducerCoroutine(numChan chan int, pExitChan chan bool) {
	for i := 0; i < 100000; i++ {
		numChan<- i
	}
	
	fmt.Println("有一个生产者协程生产完毕，退出")
	// 当前协程工作完成后存入结束标记到管道中
	pExitChan<- true
}

// 消费者
func ConsumerCoroutine(numChan, primeChan chan int, cExitChan chan bool) {
	var flag = false
	
	for {
		if num, ok := <-numChan; !ok {
			break
		} else {
			for i := 2; i < num; i++ {
				if num % i == 0 {
					flag = true
					break
				}
			}

			if flag {
				primeChan <- num
			}
		}
	}

	fmt.Println("有一个消费者协程没有数据可以消费，退出")
	// 当前协程工作完成后存入结束标记到管道中
	cExitChan<- true
}

// 单线程计算测试
func Stand() {
	start := time.Now().Unix()
	
	for num := 1; num <= 200000; num++ {
		for i := 2; i < num; i++ {
			if num % i == 0 {
				break
			}
		}
	}

	end := time.Now().Unix()
	fmt.Printf("单线程运算消耗时间: %v\n", end - start)
}

func main() {
	runtime.NumCPU()	
	
	const pNum = 2	// 生产者数量
	const cNum = 8	// 消费者数量

	// 原始数据管道
	var numChan = make(chan int, 10000)
	// 素数管道
	var primeChan = make(chan int, 20000)
	// 生产者协程结束标记管道
	var pExitChan = make(chan bool, pNum)
	// 消费者协程结束标记管道
	var cExitChan = make(chan bool, cNum)
	
	start := time.Now().Unix()
	
	// 开启生产者协程向原始数据管道存入数据
	for i := 0; i < pNum; i++ {
		go ProducerCoroutine(numChan, pExitChan)
	}

	/*
		1.开启监听协程，监听生产者标记管道的状态；
		2.若管道中存在和生产者数量一样多的标记，则代表所有生产者工作结束。
	*/
	go func() {
		for i := 0; i < pNum; i++ {
			<-pExitChan
		}
		close(numChan)
	}()

	// 开启消费者协程从原始数据管道拉取数据，判断若是素数则存入素数管道
	for i := 0; i < cNum; i++ {
		go ConsumerCoroutine(numChan, primeChan, cExitChan)
	}

	/*
		1.开启监听协程，监听消费者标记管道的状态；
		2.若管道中存在和消费者数量一样多的标记，则代表所有消费者工作结束。
	*/
	go func() {
		for i := 0; i < cNum; i++ {
			<-cExitChan
		}

		end := time.Now().Unix()
		fmt.Printf("goroutine并行计算消耗时间: %v\n", end - start)

		close(primeChan)
	}()

	for {
		if _, ok := <-primeChan; !ok {
			break
		} 
	}

	// 单线程测试
	// Stand()

	fmt.Println("主线程退出")
}