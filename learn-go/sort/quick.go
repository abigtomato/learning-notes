package main

import (
	"fmt"
)

/* 快速排序 */

func Quick(left, right int, data []int) {
	l, r := left, right	// 左右指针
	pivot := data[(left + right) / 2]	// 基准值

	// 1.for的作用就是把比pivot小的数移到左边，比pivot大的数移到右边
	for l < r {
		// 2.从左边找一个比pivot大的值
		for data[l] < pivot {
			l++
		}

		// 3.从右边找一个比pivot小的值
		for data[r] > pivot {
			r--
		}

		// 4.交换数据
		data[l], data[r] = data[r], data[l]
	}

	// 5.左右指针指向同一个数据则各走一步分开
	if l == r {
		l++
		r--
	}

	// 6.向左边递归
	if left < r {
		Quick(left, r, data)
	}

	// 7.向右边递归
	if right > l {
		Quick(l, right, data)
	}
}

func Quick001(left, right int, data []int) {
	l, r := left, right
	pivot := data[(left + right) / 2]

	for l < r {
		for data[l] < pivot {
			l++
		}

		for data[r] > pivot {
			r--
		}

		data[l], data[r] = data[r], data[l]
	}

	if l == r {
		l++
		r--
	}

	if left < r {
		Quick001(left, r, data)
	}

	if right > l {
		Quick001(l, right, data)
	}
}

func Quick002(left, right int, data []int) {
	l, r := left, right
	pivot := data[(left + right) / 2]

	for l < r {
		for data[l] < pivot {
			l++
		}

		for data[r] > pivot {
			r--
		}

		data[l], data[r] = data[r], data[l]
	}

	if l == r {
		l++
		r--
	}

	if l < right {
		Quick002(l, right, data)
	}

	if r > left {
		Quick002(left, r, data)
	}
}

func Quick003(left, right int, data []int) {
	l, r := left, right
	pivot := data[(left + right) / 2]

	for l < r {
		for data[l] < pivot {
			l++
		}

		for data[r] > pivot {
			r--
		}

		data[l], data[r] = data[r], data[l]
	}

	for l == r {
		l++
		r--
	}

	if left < r {
		Quick003(left, r, data)
	}

	if right > l {
		Quick003(l, right, data)
	}
}

func main() {
	data := []int{4, 2, 8, 0, 5, 7, 1, 3, 9}
	Quick003(0, len(data) - 1, data)
	fmt.Println(data)
}