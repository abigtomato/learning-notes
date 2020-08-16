package main

import (
	"fmt"
	"time"
)

/* 使用channel等待任务结束 */

/*
	生产者消费者模型:
	1.生产者: 发送数据端；
	2.消费者: 接收数据端；
	3.缓冲区: 
		3.1 解耦（降低生产者和消费者间的耦合度）；
		3.2 并发（生产者消费者数量不对等时，保持正常通信）；
		3.3 缓冲（生产者消费者处理速度不一致时，暂存数据）。
*/
func main() {
	// 管道做为模型中的缓冲区，无缓冲管道为模型提供同步通信，有缓冲管道则提供异步通信
	var intChan = make(chan int, 50)
	var exitChan = make(chan bool)

	// 生产者
	go func(intChan chan<- int) {
		// 生产完毕后关闭channel不再写入
		defer close(intChan)

		// intChan管道提供数据的生产消费
		for i := 0; i < 50; i++ {
			intChan <- i
			fmt.Printf("数据写入 -> %v\n", i)
			time.Sleep(time.Second)
		}
	}(intChan)

	// 消费者
	go func(intChan <-chan int, exitChan chan<- bool) {
		// 关闭标记管道让主go程判断是否消费结束
		defer close(exitChan)
		
		for {
			if v, ok := <-intChan; !ok {
				/*
					1.无缓冲管道关闭后再次读取会读出0；
					2.有缓冲管道关闭后再次读取会先读出缓冲区的数据，读完后会读出0。
				*/
				fmt.Printf("关闭后再次读取: %v\n", <-intChan)
				break
			} else {
				fmt.Printf("数据读取 -> %v\n", v)
			}
		}
	
		// exitChan管道提供消费结束的标识
		exitChan <- true
	}(intChan, exitChan)

	// 通过exitChan管道的标识判断主线程该何时结束
	for {
		// 若标记管道的写端关闭，再次读取的ok值为false
		if _, ok := <-exitChan; !ok {
			break
		}
	}
}