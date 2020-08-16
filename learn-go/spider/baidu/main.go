package main

import (
	"os"
	"fmt"
	"net/http"
	"strconv"
)

func GetPage(idx int, pageChan chan int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i - 1) * 50)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http.Get Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(fmt.Sprintf("第%d页.html", strconv.Itoa(idx)))
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()

	var ret string
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			return
		}
		if err != nil && err != io.EOFS {
			fmt.Printf("resp.Body.Read Error: %v\n", err)
			return
		}
		ret += string(buf[:n])
	}

	f.WriteString(ret)
	pageChan <- idx
}

func Working(start, end int) {
	pageChan := make(chan int)

	for i := start; i <= end; i++ {
		go GetPage(i, pageChan)
	}

	for page := range pageChan {
		fmt.Printf("第%d个页面爬取完毕\n", page)
	}
}

func main() {
	var start, end int
	fmt.Print("请输入爬取的起始页(>=1): ")
	fmt.Scan(&start)
	fmt.Print("请输入爬取的终止页(>=start): ")
	fmt.Scan(&end)

	Working(start, end)
}