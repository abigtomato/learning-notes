package main

import (
	"os"
	"fmt"
	"flag"
	"time"
	"sync"
	"io/ioutil"
	"path/filepath"
)

/* 并发目录遍历示例 */

var verbose = flag.Bool("v", false, "show verbose progress message.")	// 命令行参数
var sema = make(chan struct{}, 20)	// 限制并发数量

// 目录解析
func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			// 若是目录则新建go程递归调用本身
			go walkDir(subdir, fileSizes)
		} else {	
			// 若是文件则获取大小并入队fileSizes
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	// 通过固定缓存容量的通道来限制并发数量
	sema <- struct{}{}	// 获取令牌
	defer func() {		// 释放令牌
		<-sema
	}()
	
	// 读取目录
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	return entries
}

// 打印文件信息
func printDisUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func main() {
	// 确定初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 用于统计文件大小的通道
	fileSizes := make(chan int64)
	var wg sync.WaitGroup

	// 用于文件遍历的worker
	go func() {
		for _, root := range roots {
			wg.Add(1)
			// 解析目录
			go walkDir(root, &wg, fileSizes)
		}
		close(fileSizes)
	}()

	// 用于关闭fileSizes通道的go程
	go func() {
		wg.Wait()			// 等待所有参与遍历文件的go程工作完毕
		close(fileSizes)	// 关闭fileSizes通道以通知主go程执行下面的逻辑
	}()

	// 根据命令行参数 -v 判断是否要定时打印文件遍历信息
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			// 若通道关闭则退出select多路复用
			if !ok {
				break loop
			}

			// 统计逻辑
			nfiles++
			nbytes += size
		case <-tick:
			// 根据定时器定期打印遍历信息
			printDisUsage(nfiles, nbytes)
		}
	}

	// 信息总结
	printDisUsage(nfiles, nbytes)
}