package main

import (
	"fmt"
)

/* 归并排序 */

func Sort(data, temp[]int, left, right int) {
	// 1.左右下标未重合则继续递归
	if left < right {
		// 2.中间值
		mid := (left + right) / 2

		// 3.向左递归拆分序列
		Sort(data, temp, left, mid)
		// 4.向右递归拆分序列
		Sort(data, temp, mid + 1, right)

		// 5.合并序列
		Merge(data, temp, left, right, mid)
	}
}

func Merge(data, temp []int, left, right, mid int) {
	i, j, t := left, mid + 1, 0

	// 1.第一个for控制比较两个有序列表的值，较小的值存入temp
	for i <= mid && j <= right {
		if data[i] <= data[j] {
			temp[t] = data[i]
			t++
			i++
		} else {
			temp[t] = data[j]
			t++
			j++
		}
	}

	// 2.第二个for控制第一个有序列表的剩余值写入temp（如果存在的话）
	for i <= mid {
		temp[t] = data[i]
		t++
		i++
	}

	// 3.第三个for控制第二个有序列表的剩余值写入temp（如果存在的话）
	for j <= right {
		temp[j] = data[j]
		t++
		j++
	}

	// 4.最后一个for将temp写入源序列中
	t = 0
	tLeft := left
	for tLeft <= right {
		data[tLeft] = temp[t]
		t++
		tLeft++
	}
}

func main() {
	data := []int{4, 2, 8, 0, 5, 7, 1, 3, 9}
	Sort(0, len(data) - 1, data)
	fmt.Println(data)
}