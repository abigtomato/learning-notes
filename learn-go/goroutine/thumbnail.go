package main

import (
	"os"
	"fmt"
	"log"
	"sync"
)

/* 并行循环示例 */

func makeThumbnails(fileNames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range fileNames {
		// 为每个go程计数
		wg.Add(1)

		// 该go程为worker
		go func(f string) {
			// go程执行完逻辑后减去自己的计数
			defer wg.Done()
			
			// 具体逻辑
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}

			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// 该go程等待所有worker工作完毕，关闭size通道让主go程知道何时结束
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	// 主go程在size通道上阻塞，直到通道被关闭才会继续向下执行
	for size := range sizes {
		total += size
	}
	return total
}