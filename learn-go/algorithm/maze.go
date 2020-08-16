package other

import (
	"fmt"
	"os"
)

/*
	广度优先遍历算法找寻迷宫最短路径
 */
// 读取文件数据生产迷宫矩阵
func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	// 从文件读取指定格式的数据存入变量
	fmt.Fscanf(file, "%d %d", &row, &col)
	
	// 生成二维数组
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

// 点的封装
type point struct {
	i, j int
}

// 移动的步长
var dirs = [4]point{
	// 上左下右
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

// 移动的计算
func (this point) add(dir point) point {
	return point{this.i + dir.i, this.j + dir.j,}
}

// 根据点获取二维数组对应值
func (this point) at(grid [][]int) (int, bool) {
	// 上下越界
	if this.i < 0 || this.i >= len(grid) {
		return 0, false
	}

	// 左右越界
	if this.j < 0 || this.j >= len(grid[this.i]) {
		return 0, false
	}

	// 获取对应值
	return grid[this.i][this.j], true
}

// 开始走迷宫
func walk(maze [][]int, start, end point) [][]int {
	// steps数组记录移动步长(走出迷宫的路径上每一点需要从起点移动多少步)
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// 点队列(默认存放起始点)，每个点遍历4周相邻点之后，会根据能否通过判断是否入队(未越界，未撞墙，未走过的点入队)
	Q := []point{start}
	// 队列为空时表示没有点能够通过，迷宫无路可走
	for len(Q) > 0 {
		// 队头出队元素做为当前点
		cur := Q[0]
		// 截取除队头之外的元素做为新的队列(实现队列队头取元素的特点)
		Q = Q[1:]

		// 找到出路
		if cur == end {
			break
		}
		
		// dirs保存向4周移动的步长，遍历该数组使点向4周移动
		for _, dir := range dirs {
			// 计算当前点的其中一个邻点
			next := cur.add(dir)
			
			// 越界或撞墙
			if val, ok := next.at(maze); !ok || val == 1 {
				continue
			}
			
			// 越界或走过
			if val, ok := next.at(steps); !ok || val != 0 {
				continue
			}

			// 回到起点
			if next == start {
				continue
			}
			
			// 以上的分支全部避免则代表next点可以走
			// 在steps数组中记录到达next点的步长(当前点+1)
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			// 入队当前点可通过的下一个点(之后的某次循环会以该点为中心遍历4周，以此类推)
			Q = append(Q, next)
		}
	}

	return steps
}

// 根据步长数组生成迷宫路径
func path(end point, steps [][]int) [][]int {
	pathArr := make([][]int, len(steps))
	for i := range pathArr {
		pathArr = make([]int, len(steps[i]))
	}
	
	// 辅助变量，随循环变化，初始指向迷宫出口
	cur := end
	// 循环向上寻找入口
	for {
		curStep := cur.at(steps)
		
		// 找到入口则退出
		if curStep == 1 {
			break
		}
		
		// 遍寻4周邻点
		for dir := range dirs {
			next := cur.add(dir)
			if val, _ := next.at(steps); val == curStep - 1 {
				// 辅助变量向入口上移
				cur = next
				break
			} 
		}
	}
}

func main() {
	// 加载文件生成迷宫矩阵
	maze := readMaze("./data/maze.in")
	for _, arr := range maze {
		for _, val := range arr {
			fmt.Printf("%3v", val)
		}
		fmt.Println()
	}
	fmt.Println()

	// 走迷宫
	steps := walk(maze, point{0, 0,}, point{len(maze) - 1, len(maze[0]) - 1,})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}