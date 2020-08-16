package main

import (
	"fmt"
)

/* 冒泡排序 */

func Bubble(data []int) {
	// 1.外层循环控制行（共排序len(data) - 1次，每次循环在尾部确定一个数据的最终位置）
	for i := 0; i < len(data) - 1; i++ {
		// 2.内层循环控制列（-i的原因是序列尾部有序的数据的数量会随着外层循环的次数增加而增加）
		for j := 0; j < len(data) - 1 - i; j++ {
			// 3.相邻元素做比较
			if data[j] > data[j + 1] {
				// 4.符合规则就替换
				data[j], data[j + 1] = data[j + 1], data[j]
			}
		}
	}
}

func Bubble001(data []int) {
	for i := 0; i < len(data) - 1; i++ {
		for j := 0; j < len(data) - 1 - i; j++ {
			if data[j] > data[j + 1] {
				data[j], data[j + 1] = data[j + 1], data[j]
			}
		}
	}
}

func Bubble002(data []int) {
	for i := 0; i < len(data) - 1; i++ {
		for j := 0; j < len(data) - 1 - i; j++ {
			if data[j] > data[j + 1] {
				data[j], data[j + 1] = data[j + 1], data[j]
			}
		}
	}
}

func Bubble003(data []int) {
	for i := 0; i < len(data) - 1; i++ {
		for j := 0; j < len(data) - 1 - i; j++ {
			if data[j] > data[j + 1] {
				data[j], data[j + 1] = data[j + 1], data[j]
			}
		}
	}
}

func main() {
	data := []int{4, 2, 8, 0, 5, 7, 1, 3, 9}
	Bubble003(data)
	fmt.Println(data)
}