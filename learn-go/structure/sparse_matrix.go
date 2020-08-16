package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

/*
	[0  0  0  22  0  0  15
	 0  11 0  0   0  17 0
	 0  0  0  -6  0  0  0
	 0  0  0  0   0  39 0
	 91 0  0  0   0  0  0
	 0  0  28 0   0  0  0]
	行		列		值
	0		3		22
	0		6		15
	1		1		11
	1		5		17
	2		3		-6
	3		5		39
	4		0		91
	5		2		28 
 */

// 稀疏矩阵中的元素单位
type ValNode struct {
	Row		int			// 行
	Col 	int			// 列
	Val 	interface{}	// 值
}

// 矩阵转稀疏矩阵
func MatrixToSparseMatrix(matrix [][]interface{}) (sparseMatrix []ValNode) {
	// 预先存入稀疏矩阵的第一条数据，也就是原矩阵的行，列，默认值信息
	sparseMatrix = append(sparseMatrix, ValNode{
		Row: len(matrix),
		Col: len(matrix[0]),
		Val: 0,
	})

	for i, arr := range matrix {
		for j, val := range arr {
			// 存入非第一条数据
			if val != 0 {
				valNode := ValNode{
					Row: i,
					Col: j,
					Val: val,
				}
				sparseMatrix = append(sparseMatrix, valNode)
			}
		}
	}  

	return
}	

// 将稀疏矩阵写入文件
func WriteSparseMatrixToFile(sparseMatrix []ValNode, path string) (err error) {
	file, err := os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("os.OpenFile(path, os.O_APPEND | os.O_CREATE, 0666) fail error: %v\n", err.Error())
		return 
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, v := range sparseMatrix {
		str := fmt.Sprintf("%v\t%v\t%v\r\n", v.Row, v.Col, v.Val)
		writer.WriteString(str)
	}
	writer.Flush()

	return
}

// 从文件中读取稀疏矩阵
func ReadSparseMatrixFromFile(path string) (sparseMatrix []ValNode, err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("os.Open(path) fail error: %v\n", err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "\t")
		
		var valNode ValNode
		
		// 取出row信息
		if num, err := strconv.Atoi(arr[0]); err != nil {
			fmt.Printf("strconv.Atoi(arr[0]) fail error: %v\n", err.Error())
			break
		} else {
			valNode.Row = num
		}

		// 取出col信息
		if num, err := strconv.Atoi(arr[1]); err != nil {
			fmt.Printf("strconv.Atoi(arr[1]) fail error: %v\n", err.Error())
			break
		} else {
			valNode.Col = num
		}

		// 取出val信息
		valNode.Val = arr[2]

		sparseMatrix = append(sparseMatrix, valNode)
	}
	
	return
}

// 稀疏矩阵转矩阵
func SparseMatrixToMatrix(sparseMatrix []ValNode) (matrix [][]interface{}) {
	// 预先取出原矩阵的基本信息(也就是稀疏矩阵的第一条数据)
	valNode := sparseMatrix[0]
	// 根据基本信息初始化一个矩阵
	matrix = InitMatrix(valNode.Row, valNode.Col)
	
	for _, v := range sparseMatrix[1:] {
		matrix[v.Row][v.Col] = v.Val
	}

	return
}

// 初始化新的矩阵
func InitMatrix(oCount, iCount int) (matrix [][]interface{}) {
	for i := 0; i < oCount; i++ {
		innerArr := make([]interface{}, 0, iCount)
		for j := 0; j < iCount; j++ {
			innerArr = append(innerArr, 0)
		}
		matrix = append(matrix, innerArr)
	}

	return
}

// 遍历矩阵
func TraversingMatrix(matrix [][]interface{}) {
	for _, arr := range matrix {
		for _, val := range arr {
			fmt.Printf("%v\t", val)
		}
		fmt.Println()
	}
}

// 遍历稀疏矩阵
func TraversingSparseMatrix(sparseMatrix []ValNode) {
	for i, v := range sparseMatrix {
		fmt.Printf("sparseMatrix[%v]=%v\n", i, v)
	}
}

func main() {
	var matrix [][]interface{}
	matrix = InitMatrix(11, 11)
	matrix[0][3] = 22
	matrix[0][6] = 15
	matrix[1][1] = 11
	matrix[1][5] = 17
	matrix[2][3] = -6
	matrix[3][5] = 39
	matrix[4][0] = 91
	matrix[5][2] = 28
	fmt.Printf("type=%T, len=%v, cap=%v\n", matrix, len(matrix), cap(matrix))
	fmt.Printf("type=%T, len=%v, cap=%v\n", matrix[0], len(matrix[0]), cap(matrix[0]))
	TraversingMatrix(matrix)

	var sparseMatrix []ValNode
	sparseMatrix = MatrixToSparseMatrix(matrix)

	if err := WriteSparseMatrixToFile(sparseMatrix, "./data/SparseMatrix.model"); err != nil {
		fmt.Printf("WriteSparseMatrixToFile(sparseMatrix, \"./data/SparseMatrix.model\") fail error: %v\n", err.Error)
		return
	}

	sparseMatrix, err := ReadSparseMatrixFromFile("./data/SparseMatrix.model")
	if err != nil {
		fmt.Printf("ReadSparseMatrixFromFile(\"./data/SparseMatrix.model\") fail error: %v\n", err.Error)
		return
	}
	TraversingSparseMatrix(sparseMatrix)

	matrix = SparseMatrixToMatrix(sparseMatrix)
	TraversingMatrix(matrix)
}