package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"bufio"
)

var (
	name_re = regexp.MustCompile(`<img width="100" alt="(.*?)">`)
	score_re = regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	evaluate_re = regexp.MustCompile(`<span>(.*?)人评价</span>`)
)

func SaveToFile(idx int, nameMatchs, scoreMatchs, evaluateMatchs [][]string) {
	f, err := os.OpenFile("./data/" + strconv.Itoa(idx) + "页数据.csv", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("os.OpenFile Error: %v\n", err)
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for i := 0; i < len(nameMatchs); i++ {
		writer.WriteString(nameMatchs[i][1] + "\t" + scoreMatchs[i][1] + "\t" + evaluateMatchs[i][1] + "\r\n")
	}
	writer.Flush()
}

func ParseHtml(idx int, contents string) {
	nameMatchs := name_re.FindAllStringSubmatch(contents, -1)
	scoreMatchs := score_re.FindAllStringSubmatch(contents, -1)
	evaluateMatchs := evaluate_re.FindAllStringSubmatch(contents, -1)

	if len(nameMatchs) == len(scoreMatchs) && len(scoreMatchs) == len(evaluateMatchs) {
		SaveToFile(idx, nameMatchs, scoreMatchs, evaluateMatchs)
	} else {
		fmt.Println("抽取的信息条目数不匹配...")
		return
	}
}

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
		if n == 0 {
			return
		}
		
		if err != nil && err != io.EOF {
			fmt.Printf("resp.Body.Read Error: %v\n", err)
			return
		}

		result += string(buf[:n])
	}

	return
}

func GetPage(idx int, pageChan chan int) {
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((idx - 1) * 25) + "&filter="
	
	result, err := GetHtml(url)
	if err != nil {
		fmt.Printf("GetHtml Error: %v\n", err)
		return
	}

	ParseHtml(idx, result)

	pageChan <- idx
}

func Worker(start, end int) {
	fmt.Printf("Spider Get %d to %d Page\n", start, end)

	pageChan := make(chan int)
	for i := start; i <= end; i++ {
		go GetPage(i, pageChan)		
	}

	for idx := range pageChan {
		fmt.Printf("第 %d 页爬取完毕\n", idx)
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