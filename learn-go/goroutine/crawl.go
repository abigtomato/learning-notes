package main

import (
	"os"
	"fmt"
)

/* 并发web爬虫示例 */

func crawl(url string) []string {
	fmt.Println(url)
	
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	// 原始url通道
	worklist := make(chan []string)
	// 去重后的url通道
	unseenLinks := make(chan string)

	go func() {
		// 由命令行输入需要爬取的网页第一
		worklist <- os.Args[1:]
	}()

	// 固定20个go程参与并发
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				
				// 为防止go程在worklist阻塞，新分支一个go程处理该语句
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	// 利用map对url列表进行去重
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}