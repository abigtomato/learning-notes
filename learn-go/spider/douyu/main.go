package main

import (
	"os"
	"bufio"
	"regexp"
	"net/http"
	"fmt"
	"io"
	"strconv"
)

func GetHtml(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http.Get Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			fmt.Printf("resp.Body.Read Error: %v\n", err)
			return
		}
		result += string(buf[:n])
	}
	return
}

func SaveToImg(idx int, url string, pageChan chan int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http.Get Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	f, err := os.OpenFile("./images/" + strconv.Itoa(idx) + ".jpg", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	defer writer.Flush()

	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Printf("resp.Body.Read Error: %v\n", err)
			return
		}
		writer.Write(buf[:n])
	}

	pageChan <- idx
}

func main() {
	url := "https://www.douyu.com/g_yz"
	result, err := GetHtml(url)
	if err != nil {
		fmt.Printf("GetHtml Error: %v\n", err)
		return
	}

	pageChan := make(chan int)
	re := regexp.MustCompile(`data-original="(?s:(.*?))"`)
	matchs := re.FindAllStringSubmatch(result, -1)
	
	for idx, match := range matchs {
		go SaveToImg(idx, match[1], pageChan)
	}

	for page := range pageChan {
		fmt.Printf("第 %v 张图片爬取完成\n", page)
	}
}