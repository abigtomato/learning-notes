package main

import "fmt"

// 希尔排序
func ShellSort001(data []int) {
	// 第一层for控制增量递减，inc表示当前增量(步长)
	for inc := len(data) / 2; inc > 0; inc /= 2 {
		// 第二层for控制步长后方数据的变化
		for i := inc; i < len(data); i++ {
			temp := data[i]

			// 第三层for控制步长前方数据的变化
			for j := i - inc; j >= 0; j -= inc {
				// 交互数据
				if temp < data[j] {
					data[j], data[j + inc] = data[j + inc], data[j]
				} else {
					break
				}
			}
		}
	}
}

func ShellSort002(data []int) {
	var temp int
	for inc := len(data) / 2; inc > 0; inc /= 2 {
		for i := inc; i < len(data); i++ {
			temp = data[i]
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
	data := []int{2, 1, 6, 8, 3, 5, 9, 4, 7}
	ShellSort001(data)
	ShellSort002(data)
	fmt.Println(data)
}
