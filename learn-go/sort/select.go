package main

import (
	"fmt"
)

/* 选择排序 */

func Select(data []int) {
	// 1.外层遍历整个序列
	for i := 0; i < len(data); i++ {
		// 2.默认0号元素为最大值
		max := 0

		// 3.遍历剩下的元素，找出最大值
		for j := 1; j < len(data) - i; j++ {
			if data[j] > data[max] {
				max = j
			}
		}

		// 4.将最大值移动到序列后面（序列末尾确定一个元素位置后，下次遍历就可以忽略）
		data[max], data[len(data) - 1 - i] = data[len(data) - 1 -i], data[max]
	}
}

func Select001(data []int) {
	for i := 0; i < len(data); i++ {
		max := 0

		for j := 1; j < len(data) - i; j++ {
			if data[j] > data[max] {
				max = j
			}
		}

		data[max], data[len(data) - 1 - i] = data[len(data) - 1 - i], data[max]
	}
}

func Select002(data []int) {
	for i := 0; i < len(data); i++ {
		max := 0

		for j := 1; j < len(data) - i; j++ {
			if data[j] > data[max] {
				max = j
			}
		}

		data[max], data[len(data) - 1 - i] = data[len(data) - 1 - i], data[max]
	}
}

func Select003(data []int) {
	for i := 0; i < len(data); i++ {
		max := 0

		for j := 1; j < len(data) - i; j++ {
			if data[j] > data[max] {
				max = j 
			}
		}

		data[max], data[len(data) - 1 - i] = data[len(data) - 1 - i], data[max]
	}
}

func main() {
	data := []int{4, 2, 8, 0, 5, 7, 1, 3, 9}
	Select003(data)
	fmt.Println(data)
}