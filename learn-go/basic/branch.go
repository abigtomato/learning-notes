package main

import (
	"fmt"
	"io/ioutil"
)

/* 分支结构 */

// 1.if分支
func opt() {
	const finename = "./data/demo.txt"
	// if的判断条件中可以先定义变量并赋值；
	// ioutil库提供io流常用的函数；
	// ReadFile提供读取文件内容的功能，返回值为[]byte类型的文件内容和错误信息。
	if contents, err := ioutil.ReadFile(finename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(contents))
	}
}

// 2.switch分支
func grade(score int) string {
	var g string
	
	switch {
	// case中的break由go编译器自动添加
	case score < 0 || score > 100:
		// panic提供报错机制
		// fmt.Sprintf返回格式化后的字符串
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60, score == 60, score <= 60:
		g = "D"
	case score < 80:
		g = "C"
		// fallthrough关键字默认穿透一层case
		fallthrough	
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	// 所有case条件都不符合，执行default默认
	default:
		fmt.Println("Default ......")
	}
	
	return g
}

func main() {
	opt()

	fmt.Println(grade(50), grade(70), grade(80), grade(90))
}