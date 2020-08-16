package main

import (
	"fmt"
)

/* 插入排序 */

func Insert(data []int) {
	// 1.外层for遍历无序表，i表示无序表第一个元素（初始时：0为有序表，1~len-1为无序表）
	for i := 1; i < len(data); i++ {
		// 2.j表示无序表第一个元素的前一个元素，也就是有序表最后一个元素（有序表无序表连在一起）
		j := i - 1
		// 3.每次循环从无序表取出的第一个元素，也就是本次循环需要插入有序表的新元素
		temp := data[i]

		// 4.新元素从有序表尾部开始循环比较，直到找到适合的插入位置
		for j >= 0 && data[j] > temp {
			data[j + 1] = data[j]	// 为新元素的插入腾出位置
			j--	// 无序表中向前移动
		}

		// 5.插入到有序表
		data[j + 1] = temp
	}
}

func Insert001(data []int) {
	for i := 1; i < len(data); i++ {
		j := i - 1
		temp := data[i]

		for j >= 0 && data[j] > temp {
			data[j + 1] = data[j]
			j--
		}

		data[j + 1] = temp
	}
}

func Insert002(data []int) {
	for i := 1; i < len(data); i++ {
		j := i - 1
		temp := data[i]

		for j >= 0 && data[j] > temp {
			data[j + 1] = data[j]
			j--
		}

		data[j + 1] = temp
	}
}

func Insert003(data []int) {
	for i := 1; i < len(data); i++ {
		j := i - 1
		temp := data[i]

		for j >= 0 && data[j] > temp {
			data[j + 1] = data[j]
			j--
		}

		data[j + 1] = temp
	}
}

func main() {
	data := []int{4, 2, 8, 0, 5, 7, 1, 3, 9}
	Insert003(data)
	fmt.Println(data)
}