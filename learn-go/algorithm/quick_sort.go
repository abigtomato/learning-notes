package main

import "fmt"

/*
	快速排序

	初始状态:
		l = 0, r = 5, pivot = 2
		[-9, 78, 0, 23, -567, 7]
	第一次排序:
		l = 1, r = 4, arr[l] = 78, arr[r] = -567
		[-9, -567, 0, 23, 78, 7]
		当l = 3, r = 1, arr[l] = 23, arr[r] = -567时，l >= r 退出第一次排序
	左递归:
		QuickSort(left, r, arr)
		left = 0, r = 1, arr = [-9, -567]
	右递归:
		QuickSort(l, right, arr)
		l = 3, right = 5, arr = [23, 78, 7]
	分析:
		1. 通过第一次排序将要排序的数据按照中间值分割成独立的两部分
		2. 期望的情况是左边部分的所有数据都比中间值要小，右边部分要比中间大
		3. 之后再次按照此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行
 */
func QuickSort001(left int, right int, arr []int) {
	// 左右指针 := 起始结束指针
	l, r := left, right
	// 中间值
	pivot := arr[(left + right) / 2]

	// 该循环控制左右指针的移动
	for l < r {
		// 从左边查找大于中间值的元素下标
		for arr[l] < pivot {
			l++
		}

		// 从右边查找小于中间值的元素下标
		for arr[r] > pivot {
			r--
		}

		// 当指针重合或错位则退出循环
		if l >= r {
			break
		}

		// 交换数据
		arr[l], arr[r] = arr[r], arr[l]

		// 若出现和中间值相等的情况则无视，继续移动指针
		if arr[l] == pivot {
			l++
		}

		// 若出现和中间值相等的情况则无视，继续移动指针
		if arr[r] == pivot {
			r--
		}
	}

	// 若指针重合，则各自向自己移动的方向移动一步
	if l == r {
		l++
		r--
	}

	// 1. 若起始指针小于右指针，证明还有空间可以划分
	// 2. 以起始指针到右指针的范围单独看成一段空间，递归按照快排的逻辑处理
	if left < r {
		QuickSort001(left, r, arr)
	}

	// 1. 若结束指针大于左指针，证明还有空间可以划分
	// 2. 以左指针到结束指针的范围单独看成一段空间，递归按照快排的逻辑处理
	if right > l {
		QuickSort001(l, right, arr)
	}
}

func QuickSort002(left int, right int, data []int) {
	l, r := left, right
	pivot := data[(left + right) / 2]

	for l < r {
		for data[l] < pivot {
			l++
		}

		for data[r] > pivot {
			r--
		}

		if l >= r {
			break
		}

		data[l], data[r] = data[r], data[l]

		if data[l] == pivot {
			l++
		}

		if data[l] == pivot {
			r--
		}
	}

	if l == r {
		l++
		r--
	}

	if left < r {
		QuickSort002(left, r, data)
	}

	if right > l {
		QuickSort002(l, right, data)
	}
}

func main() {
	data := []int{2, 1, 6, 8, 3, 5, 9, 4, 7}
	QuickSort001(0, len(data) - 1, data)
	QuickSort002(0, len(data) - 1, data)
	fmt.Println(data)
}
