package main

import (
	"fmt"
	"errors"
)

/* 
	recover()捕获异常
 */

// 使用defer + recover()来捕获处理异常
func tryRecover() {
	// defer将匿名函数调用语句压入defer栈中，等待函数执行结束调用
	defer func() {
		// recover()内置函数可以捕获到异常
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Printf("Error occurred: %s\n", err.Error())
		} else {
			panic(r)
		}
	}()

	num1 := 10
	num2 := 0
	fmt.Println("result: %v\n", num1 / num2)
}

// 抛出自定义错误测试
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("配置文件读取错误......")
	}
}

// 测试自定义错误
func errorTest2() {
	if err := readConf("config2.ini"); err != nil {
		// 若出错，则打印自定义错误的信息，终止程序
		panic(err)
	}
}

func main() {
	tryRecover()
	errorTest2()
}