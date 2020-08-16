package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
	"math/rand"
	"time"
)

/* 循环结构 */

// 1.for循环（10进制转2进制例子）
func convertToBin(n int) string {
	var result string

	// 无起始式，有迭代式形式的for循环
	for ; n > 0; n /= 2 {
		result = strconv.Itoa(n % 2) + result
	}

	return result
}

// 2.for循环（无起始和迭代式，按行读取文件）
func printFile(filename string) {
	// os包提供操作系统相关的函数库
	// Open()提供打开文件的功能
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// bufio包提供缓冲流的函数库
	// NewScanner()按照文件创建扫描器
	scanner := bufio.NewScanner(file)
	// Scan()每次扫描文件一行并向后移动行标记，扫描到末尾则返回nil
	for scanner.Scan() {
		// Text()将当前标记指向的行转换为文本返回
		fmt.Println(scanner.Text())
	}
}

// 3.for循环模拟while
func for2while() {
	// 使用Unix时间戳做为随机数种子，使产生的随机数不会重复
	rand.Seed(time.Now().Unix())
	
	j := 1
	for j <= 10 {
		// 生成[0, 100)间的整数
		// +1 是为了生成[1, 100]的整数
		n := rand.Intn(100) + 1
		fmt.Printf("Docker/Kubernetes [%d]\n", n)
		j++
	}

	k := 1
	for ; ; {
		if k > 10 {
			break
		}
		k++
	}
	fmt.Printf("k = %d\n", k)
}

// 4.死循环
func forever() {
	for {
		fmt.Println("Golang")
	}
}

// 5.打印空心金字塔案例
func pyramid(totalLevel int) {
	// 层数控制
	for i := 1; i <= totalLevel; i++ {
		// 打印空格 => 空格的规律: 总层数-当前层数
		for j := 1; j <= totalLevel - i; j++ {
			fmt.Print(" ")
		}
		// 打印* => *的规律: 2*当前层数-1
		for k := 1; k <= 2 * i - 1; k++ {
			// 控制中间空出
			if k == 1 || k == 2 * i - 1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println(convertToBin(5), convertToBin(13))
	fmt.Println()

	printFile("./data/demo.txt")
	fmt.Println()

	for2while()
	fmt.Println()

	pyramid(9)
	fmt.Println()

	// forever()
	
	// goto跳转到指定标记处执行代码
	var num int = 30
	if num > 20 {
		goto label1
	}
	
	fmt.Println("go ...")
	label1:
	fmt.Println("goto ...")
}