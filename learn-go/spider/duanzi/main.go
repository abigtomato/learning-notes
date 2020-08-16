package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"net/http"
	"strconv"
	"strings"
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
		if err != nil && err != io.EOF {
			fmt.Printf("resp.Body.Read Error: %v\n", err)
			return
		}
		result += string(buf[:n])
	}
	return
}

func SaveToFie(idx int, infoMap map[string]string) (err error) {
	file, err := os.OpenFile("./data/" + strconv.Itoa(idx) + ".csv", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("os.OpenFile Error: %v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	buf := make([]byte, 4096)
	for k, v := range infoMap {
		writer.WriteString(k + "\t" + v + "\r\n")
	}

	return
}

func SpiderJoke(idx int, url string) (title, content string) {
	result, err := GetHtml(url)
	if err != nil {
		fmt.Printf("GetHtml Error: %v\n", err)
		return
	}

	title_re := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	title_matchs := title_re.FindAllStringSubmatch(result, 1)
	for _, match := range title_matchs {
		title = strings.Replace(match[1], "\t", "", -1)
		break
	}

	content_re := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	content_matchs := content_re.FindAllStringSubmatch(result, -1)
	for _, match := range content_matchs {
		content = strings.Replace(match[1], "\n", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		break
	}

	return
}

func SpiderPage(idx int, url string, pageChan chan int) {
	result, err := GetHtml(url)
	if err != nil {
		fmt.Printf("GetHtml Error: %v\n", err)
		return
	}

	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	matchs := re.FindAllStringSubmatch(result, -1)

	infoMap := make(map[string]string)
	for _, match := range matchs {
		title, content := SpiderJoke(idx, match[1])
		infoMap[title] = content
	}

	if err := SaveToFie(idx, infoMap); err != nil {
		fmt.Printf("SaveToFile Error: %v\n", err)
		return
	}

	pageChan <- idx
}

func Worker(start, end int) {
	fmt.Printf("Spider %d to %d Page\n", start, end)

	pageChan := make(chan int)
	for i := start; i <= end; i++ {
		url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
		go SpiderPage(i, url, pageChan)
	}

	for idx := range pageChan {
		fmt.Printf("第 %v 页爬取完毕\n", idx)
	}
}

func main() {
	var start, end int
	fmt.Print("Start: ")
	fmt.Scan(&start)
	fmt.Print("End: ")
	fmt.Scan(&end)

	Worker(start, end)
}