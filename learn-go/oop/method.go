package main

import "fmt"

// 节点结构体
type TreeNode struct {
	Value int	// 节点值
	Left, Right *TreeNode	// 指向左右子树的指针
}

// 将函数绑定到指定类型上，该函数就是指定类型的方法
func (node TreeNode) Print() {
	// 若是值接收者会存在新的拷贝，无法改变原结构体的内容
	fmt.Printf("%v\t", node.Value)
}

func (node *TreeNode) SetValue(value int) {
	// 指针接收者才能改变原结构体内容
	(*node).Value = value
}

// 前序遍历
func (node *TreeNode) Prologue() {
	if node == nil {
		return
	}

	(*node).Print()
	(*node).Left.Prologue()
	(*node).Right.Prologue()
}

// 中序遍历
func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}

	(*node).Left.Traverse()
	(*node).Print()
	(*node).Right.Traverse()
}

// 后序遍历
func (node *TreeNode) PostOrder() {
	if node == nil {
		return
	}

	// 编译器会将node转换为(*node)的形式调用
	node.Left.PostOrder()
	node.Right.PostOrder()
	node.Print()
}

// 绑定String()方法，可以打印出实例的信息
func (node *TreeNode) String() string {
	return fmt.Sprintf("Value=[%v], Left=[%v], Right=[%v]\n", node.Value, node.Left, node.Right)
}

// 工厂函数
func CreateNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

// 为int类型起别名
type integer int

// 自定义数据类型都可以存在方法
func (i integer) print() {
	fmt.Printf("i=%v\n", i)
}

func (i *integer) change() {
	*i = *i + 1
}

func main() {
	var root TreeNode = TreeNode{Value: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{4, nil, nil}
	root.Right = &TreeNode{5, nil, nil}
	root.Left.Right = CreateNode(6)
	root.Right.Left = new(TreeNode)
	fmt.Println(&root)
	fmt.Println()

	// 方法调用和传参机制分析:
	// 1.方法调用后会在栈区分配专属的方法栈
	// 2.若是值传递，调用时会将调用者(结构体)同参数一并传入方法的栈中，相当于方法栈中存在一份调用者的拷贝
	// 3.若是引用传递，方法栈中会存在指针指向调用者
	(&root).Prologue()
	fmt.Println()

	(&root).Traverse()
	fmt.Println()

	// 当指针传递时，编译器底层会将root转换为&root传递
	root.PostOrder()
	fmt.Println()

	// 测试为自定义类型绑定方法
	var i integer = 10
	i.change()
	i.print()
}