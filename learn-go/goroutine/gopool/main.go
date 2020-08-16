package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"go.learn/go.pool/pool"
)

func main() {
	var wg sync.WaitGroup
	defer pool.Close()

	runTimes := 100000
	for i := 0; i < runTimes; i++ {
		wg.Add(1)	// go程加锁
		if err := pool.Go(func() {
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Hello Goroutine Pool!") // 任务逻辑
			wg.Done()	// go程解锁
		}); err != nil {
			log.Println(err.Error())
		}
	}
	wg.Wait()	// 主go程等待
}
