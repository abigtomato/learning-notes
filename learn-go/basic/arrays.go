package main

import (
	"fmt"
	"time"
	"math/rand"
)

/* 数组 */

// 数组值传递（[5]int 表示接收长度为5类型为int的数组）
func PrintArray1(arr [5]int) {
	arr[0] = 100
	// 迭代返回下标和值
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 数组指针传递（*[5]int 表示接收长度为5类型为int的数组的内存地址）
func PrintArray2(arr *[5]int)  {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 数组的遍历
func Traversing(arr [5]int, grid [4][5]int) {
	// 遍历一维数组
	for i, v := range arr {
		fmt.Printf("(index=%v, value=%v)", i, v)
	}
	fmt.Println()

	// 遍历二维数组
	for i := range grid {
		// 若无需第一个返回值，可用_代替
		for _, v := range grid[i] {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
}

// 示例1：利用编码的顺序，让字符参与运算存储26个大写字母到数组中
func ArrayDemo1() {
	var myChars [26]byte
		
	for i := 0; i < len(myChars); i++ {
		myChars[i] = 'A' + byte(i)
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c", myChars[i])
	}

	fmt.Println()
}

// 示例2：随机生成数组，反转打印
func ArrayDemo2() {
	var intArr [5]int
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println("交换前: ", intArr)
	
	for i := 0; i < len(intArr) / 2; i++ {
		intArr[i], intArr[len(intArr) - 1 - i] = intArr[len(intArr) - 1 - i], intArr[i]
	}
	fmt.Println("交换后: ", intArr)
}

func main() {
	// 1.数组创建
	var arr1 [3]int = [3]int{0: 1, 1: 2, 2: 3}	// 显式声明（指定类型和下标）
	var arr2 = [3]int{1, 2, 3}	// 声明并预创建初值
	arr3 := [...]int{1, 2, 3, 4, 5}	// [...]声明不定长数组
	var grid [4][5]int	// 二维数组，值都默认为0
	fmt.Println(arr1, arr2, arr3, grid)

	// 数组的遍历
	Traversing(arr3, grid)

	// 若是值传递，则会在函数中拷贝出一份新数组
	fmt.Println(arr3)
	PrintArray1(arr3)
	fmt.Println(arr3)

	// 若是指针传递，则在函数中操作则会改变原数组
	fmt.Println(arr3)
	PrintArray2(&arr3)
	fmt.Println(arr3)

	// 测试案例
	ArrayDemo1()
	ArrayDemo2()
	fmt.Println()

	// 数组的内存布局分析:
	// 	1.数组是值类型，数组的地址就是数组首元素的地址；
	// 	2.数组是一段连续的内存空间，每个元素的间隔是由元素大小决定的。
	var arr [3]int
	fmt.Printf("数组的地址=%p, 首元素的地址=%p, 第二个元素的地址=%p, 第三个元素的地址=%p\n",
		&arr, &arr[0], &arr[1], &arr[2])	// %p格式化输出地址
	fmt.Println()

	// 二维数组内存布局分析:
	// 	1.存储多个指针，分别指向底层的一维数组地址；
	// 	2.数组在内存中的空间是连续的。
	var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("数据=%v, 类型=%T, 二维数组地址=%p, 二维数组第一个指针指向=%p, 二维第二个指针指向=%p\n",
		matrix, matrix, &matrix, &matrix[0], &matrix[1])
	fmt.Printf("第一个一维数组地址=%p, 第二个一维数组地址=%p\n", &matrix[0][0], &matrix[1][0])	
	fmt.Println()
}