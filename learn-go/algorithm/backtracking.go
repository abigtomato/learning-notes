package other

import "fmt"

/*
	递归实现迷宫回溯
	规则:
		0表示路径
		1表示墙
		2表示通过的路径
 		3表示死路
 */
func SetWay(myMap *[8][7]int, i, j int) bool {
	// 预先设定myMap[6][5]点为终点，若终点为2表示已通过，则迷宫通过
	if myMap[6][5] == 2 {
		return true
	} else {
		// 若当前位于的点为0，表示是路径，可以通过
		if myMap[i][j] == 0 {
			// 预先设定该位置是2，通过的路径
			myMap[i][j] = 2
			// 向下，右，上，左方向递归求解
			if SetWay(myMap, i + 1, j) {
				return true
			} else if SetWay(myMap, i, j + 1) {
				return true
			} else if SetWay(myMap, i - 1, j) {
				return true
			} else if SetWay(myMap, i, j - 1) {
				return true
			} else {
				// 若4种方向的递归收敛后返回的都是false，则代表该位置不通是死路
				myMap[i][j] = 3
				return false
			}
		} else {
			// 若当前位置不是0路径，退出
			return false
		}
	}
}

func main() {
	var myMap [8][7]int

	for i := 0; i < len(myMap) - 1; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}

	for i := 0; i < len(myMap); i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	myMap[3][1] = 1
	myMap[3][2] = 1

	for _, arr := range myMap {
		for _, val := range arr {
			fmt.Printf("%v ", val)
		}
		fmt.Println()
	}

	SetWay(&myMap, 1, 1)

	fmt.Println()
	for _, arr := range myMap {
		for _, val := range arr {
			fmt.Printf("%v ", val)
		}
		fmt.Println()
	}
}