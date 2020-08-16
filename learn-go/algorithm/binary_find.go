package main

import "fmt"

// 二分查找
func BinaryFind001(data []int, left int, right int, value int) {
	// 左下标超过右下标时递归结束，表示无法找到
	if left > right {
		return
	}

	// 中间下标，将查找区间分为前后两部分
	middle := (left + right) / 2

	// 若是查找的数大于中间的数据，那么将左下标移动到中间下标的后一位，缩短最大查找范围为后半部分
	if value > data[middle] {
		BinaryFind001(data, middle + 1, right, value)
	} else if value < data[middle] {	// 若是查找的数小于中间的数据，那么将右下标移动到中间下标的前一位，缩短最大查找范围为前半部分
		BinaryFind001(data, left, middle - 1, value)
	} else {
		fmt.Println(middle)
	}
}

func BinaryFind002(data []int, left int, right int, value int) {
	if left > right {
		return
	}

	middle := (left + right) / 2

	if value > data[middle] {
		BinaryFind002(data, middle + 1, right, value)
	} else if value < data[middle] {
		BinaryFind002(data, left, middle - 1, value)
	} else {
		fmt.Println(middle)
	}
}

func main() {
	data := []int{1, 8, 10, 89, 1000, 1234}
	BinaryFind001(data, 0, len(data) - 1, 1000)
	BinaryFind002(data, 0, len(data) - 1, 1000)
}
