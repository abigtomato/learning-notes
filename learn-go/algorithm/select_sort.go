package main

import "fmt"

/*
	选择排序

	初始状态:
		[8, 3, 2, 1, 7, 4, 6, 5]
	第1次比较并交换:
		[1, 3, 2, 8, 7, 4, 6, 5]
	第2次比较并交换:
		[1, 2, 3, 8, 7, 4, 6, 5]
	第3次比较并交换:
		[1, 2, 3, 4, 7, 8, 6, 5]
	第4次比较并交换:
		[1, 2, 3, 4, 5, 8, 6, 7]
	第5次比较并交换:
		[1, 2, 3, 4, 5, 6, 8, 7]
	第6次比较并交换:
		[1, 2, 3, 4, 5, 6, 7, 8]
	第7次比较并交换:
		[1, 2, 3, 4, 5, 6, 7, 8]
	分析 (升序):
		1. 第1次从arr[0] ~ arr[n-1]中选取最小值，与arr[0]交换
		2. 第2次从arr[1] ~ arr[n-1]中选取最小值，与arr[1]交换
		3. 第3次从arr[2] ~ arr[n-1]中选取最小值，与arr[2]交换
		4. 第i次从arr[i-1] ~ arr[n-1]中选取最小值，与arr[i-1]交换
		5. 第n-1次从arr[n-2] ~ arr[n-1]中选取最小值，与arr[n-2]交换
		6. 总共通过n-1次，得到一个按排序码从小到大排序的有序序列
 */
func SelectSort001(data []int) {
	// 外层控制行
	for i := 0; i < len(data); i++ {
		// 记录最大值用的下标(默认置为第一个)
		index := 0

		// 1. 内层找寻除尾部稳定位的其他位的最大数据
		// 2. 每次内层for都会找寻最大值移动到尾部合适位置
		// 3. j从1开始是因为0位默认为最大值
		for j := 1; j < len(data) - i; j++ {
			if data[j] > data[index] {
				index = j
			}
		}

		// 将本次最大值和"最后一个元素"(不包含位置固定的元素)交换
		data[index], data[len(data) - 1 - i] = data[len(data) - 1 - i], data[index]
	}
}

func SelectSort002(data []int) {
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
	data := []int{2, 1, 6, 8, 3, 5, 9, 4, 7}
	SelectSort001(data)
	SelectSort002(data)
	fmt.Println(data)
}
