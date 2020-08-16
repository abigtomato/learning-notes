package main

import (
	"fmt"
	"time"
	"math/rand"
)

/* 使用select + channel生产消费斐波那契数列 */

// 创建生产者
func Producer(pid, total int) chan int {
	pChan := make(chan int)
	
	go func() {
		defer close(pChan)	

		x, y := 1, 1
		for i := 0 ; i < total; i++ {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			fmt.Printf("生产者%d号 -> %d\n", pid, x)
			
			pChan <- x
			x, y = y, x + y
		}
	}()

	return pChan
}

// 创建工作go程
func CreateWorker() (chan int, chan bool) {
	// 工作go程有自己的数据管道和退出管道
	intChan := make(chan int, 3)
	quitChan := make(chan bool)
	
	go func() {
		defer func() {
			close(intChan)
			close(quitChan)
		}()

		for num := range intChan {
			fmt.Printf("Worker Received: %d\n", num)
		}
		quitChan <- true
	}()

	return intChan, quitChan
}

func main() {
	// 原始数据管道
	pChan1, pChan2 := Producer(1, 20), Producer(2, 30)

	// 定时器
	tm := time.After(100 * time.Second)	// 间隔指定时间生产数据
	tick := time.Tick(time.Second)	// 返回通道，每隔一段指定时间生产数据

	// worker实例
	worker, quitChan := CreateWorker()
	
	// 待消费队列
	var values []int
	
	for {
		// 当前正在工作的worker
		var activeWorker chan int
		// 当前需要被消费的队头数据
		var activeValue int
		
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		
		/*
			select语句（监听管道的数据流动）：
			1.按照顺序从头到尾评估每一个case后面的IO操作；
			2.当任意一个case可执行（即管道解阻塞），则该case会执行；
			3.若本次有多个case可以执行，那么从可执行的case中任意选择一条执行；
			4.若本次没有任意一条语句可以执行（即所有case的管道都阻塞），则:
				4.1 若存在default语句，则本次select执行default；
				4.2 若无default，那么select会阻塞，直到至少有一个通信可以进行下去。
		*/
		select {
			case num := <-pChan1:
				values = append(values, num)
			case num := <-pChan2:
				values = append(values, num)
			case activeWorker <-activeValue:
				values = values[1:]
			case <-tick:
				fmt.Printf("队列长度: %v\n", len(values))
			case <-time.After(1500 * time.Millisecond):
				fmt.Println("select监听超时")
			case <-tm:
				return
			case <-quitChan:
				return
		}
	}
}