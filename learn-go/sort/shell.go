package main

import (
	"fmt"
)

/* 希尔排序 */

func Shell(data []int) {
	// 1.外层for控制步长的增量
	for inc := len(data) / 2; inc > 0; inc /= 2 {
		// 2.第2层for控制步长后端元素的后移
		for i := inc; i < len(data); i++ {
			// 3.临时存储步长后端元素
			temp := data[i]
			
			// 4.内存for控制步长前端元素的后移
			for j := i - inc; j >= 0; j -= inc {
				// 5.比较步长两端的元素，满足条件互换
				if temp < data[j] {
					data[j], data[j + inc] = data[j + inc], data[j]
				} else {
					break
				}
			}
		}
	}
}

func Shell001(data []int) {
	for inc := len(data) / 2; inc > 0; inc /= 2 {
		for i := inc; i < len(data); i++ {
			temp := data[i]

			for j := i - inc; j >= 0; j -= inc {
				if temp < data[j] {
					data[j], data[j + inc] = data[j + inc], data[j]
				} else {
					break
				}
			}
		}
	}
}

func Shell002(data []int) {
	for inc := len(data) / 2; inc > 0; inc /= 2 {
		for i := inc; i < len(data); i++ {
			temp := data[i]

			for j := i - inc; j >= 0; j -= inc {
				if data[j] > temp {
					data[j], data[j + inc] = data[j + inc], data[j]
				} else {
					break
				}
			}
		}
	}
}

func Shell003(data []int) {
	for inc := len(data) / 2; inc > 0; inc /= 2 {
		for i := inc; i < len(dat); i++ {
			temp := data[i]

			for j := i - inc; j >= 0; j -= inc {
				if temp < data[j] {
					data[j], data[j + inc] = data[j + inc], data[j]
				} else {
					break
				}
			}
		}
	}
}
 
func main() {
	data := []int{4, 2, 8, 0, 5, 7, 1, 3, 9}
	Shell003(data)
	fmt.Println(data)
}