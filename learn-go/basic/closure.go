package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)

/* 
	闭包
 */

// 例1：累加器
func AddUpper() func(int) int {
	sum := 0
	// 匿名函数和自己引用的外部环境(AddUpper()内部的num变量就是外部环境)形成的整体，称之为闭包
	return func(num int) int {
		sum += num
		return sum
	}
}

// 例2：添加后缀
func MakeSuffix(suffix string) func(string) string {
	return func(name string) string {
		// 这里的suffix就是外部引用资源
		if !strings.HasSuffix(name, suffix) {
			return name + "." + suffix
		}
		return name
	}
}

// 例3：斐波那契数列
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

// 例4：函数实现接口
type intGen func() int	// 自定义类型别名，只要是自定义类型都能实现接口

// 实现Read方法从而实现io.Read接口
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

// scan读取实现了io.Read接口的类型
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 例5：实现二叉树节点遍历
type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

func (this *TreeNode) Print() {
	fmt.Printf("[%v] ", this.Value)
}

func (this *TreeNode) TraverseFunc(fun func(*TreeNode)) {
	if this == nil {
		return
	}

	this.Left.TraverseFunc(fun)
	fun(this)
	this.Right.TraverseFunc(fun)
}

func main() {
	// 例1：累加器
	// AddUpper()返回一个闭包，闭包类比于class，闭包中的函数是类的方法，引用的外部环境是类的属性
	add := AddUpper()
	// 反复调用闭包，闭包中的外部环境资源只会初始化一次，多次调用就成了累加
	for i := 0; i < 10; i++ {
		fmt.Printf("1 + 2 + ... + %v = %v\n", i, add(i))
	}
	fmt.Println()

	// 例2：添加后缀
	// 这里返回一个为字符串添加后缀的闭包
	makeSuffix := MakeSuffix("jpg")
	fmt.Println(makeSuffix("Hadoop.jpg"))
	fmt.Println(makeSuffix("Spark"))
	fmt.Println(makeSuffix("Storm"))
	fmt.Println()

	// 例3：斐波那契数列
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
	fmt.Println()

	// 例4：函数实现接口
	printFileContents(fib)
	fmt.Println()

	// 例5：实现二叉树节点遍历
	root := &TreeNode{
		Value: 3,
		Left: &TreeNode{
			Value: 0,
			Right: &TreeNode{Value: 2,},
		},
		Right: &TreeNode{
			Value: 5,
			Left: &TreeNode{Value: 4,},
		},
	}

	fmt.Printf("中序遍历: ")
	root.TraverseFunc(func(node *TreeNode) {
		node.Print()
	})
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *TreeNode) {
		nodeCount++
	})
	fmt.Printf("二叉树节点数量: %v\n", nodeCount)
}