package main

import (
	"fmt"
	"sync"
)

/* 使用sync.WaitGroup等待任务结束 */

// worker结构
type worker struct {
	in 		chan int
	done 	func()
}

// 创建worker
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			// 解除waitgroup中的一个go程锁定
			wg.Done()
		},
	}
	
	// 开启消费协程，从worker.in消费数据
	go func(id int, w worker) {
		for n := range w.in {
			fmt.Printf("Worker %d Received %c\n", id, n)
			// 消费完毕解除锁定
			w.done()
		}
	}(id, w)
	
	return w
}

func master() {
	var wg sync.WaitGroup

	// 启动10个工作协程
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
		// 往waitgroup中注册协程任务
		wg.Add(1)
	}

	// 生产数据到worker.in
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	// 使用waitgroup使主go程等待
	wg.Wait()
}

func main() {
	master()
}