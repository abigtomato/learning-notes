package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Data 	interface{}
	Left	*TreeNode 
	Right	*TreeNode
}

func NewTree() *TreeNode {
	root := new(TreeNode)
	return root
}

func (node *TreeNode) PreOrder() {
	if node == nil {
		return
	}

	fmt.Printf("%v ", node.Data)
	node.Left.PreOrder()
	node.Right.PreOrder()
}

func (node *TreeNode) MidOrder() {
	if node == nil {
		return
	}

	node.Left.MidOrder()
	fmt.Printf("%v ", node.Data)
	node.Right.MidOrder()
}

func (node *TreeNode) RearOrder() {
	if node == nil {
		return
	}

	node.Left.RearOrder()
	node.Right.RearOrder()
	fmt.Printf("%v ", node.Data)
}

func (node *TreeNode) Height() int {
	if node == nil {
		return 0
	}

	lh := node.Left.Height()
	rh := node.Right.Height()

	if lh > rh {
		lh++
		return lh
	} else {
		rh++
		return rh
	}
}

func (node *TreeNode) LeafCount(count *int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		(*count)++
	}

	node.Left.LeafCount(count)
	node.Right.LeafCount(count)
}

func (node *TreeNode) Search(data interface{}) {
	if node == nil {
		return
	}

	if reflect.TypeOf(node.Data) == reflect.TypeOf(data) && reflect.DeepEqual(node.Data, data) {
		fmt.Printf("数据已找到: %v\n", data)
		return
	}

	node.Left.Search(data)
	node.Right.Search(data)
}

func (node *TreeNode) Destroy() {
	if node == nil {
		return
	}

	node.Left.Destroy()
	node.Left = nil
	node.Right.Destroy()
	node.Right = nil	

	node.Data = nil
}	

func (node *TreeNode) Reverse() {
	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left

	node.Left.Reverse()
	node.Right.Reverse()
}

func (node *TreeNode) Copy() *TreeNode {
	if node == nil {
		return nil
	}

	left := node.Left.Copy()
	right := node.Right.Copy()

	new := new(TreeNode)
	new.Data = node.Data
	new.Left = left
	new.Right = right
	
	return new
}

func main() {
	root := NewTree()
	root.Data = 0
	root.Left = &TreeNode{
		Data: 1,
		Left: &TreeNode{
			Data: 3,
		},
		Right: &TreeNode{
			Data: 4,
		},
	}
	root.Right = &TreeNode{
		Data: 2,
		Left: &TreeNode{
			Data: 5,
		},
		Right: &TreeNode{
			Data: 6,
		},
	}

	root.PreOrder()
	fmt.Println()

	root.MidOrder()
	fmt.Println()
	
	root.RearOrder()
	fmt.Println()

	fmt.Println(root.Height())

	var count int
	root.LeafCount(&count)
	fmt.Println(count)

	root.Search(6)

	//root.Destroy()
	
	root.PreOrder()
	fmt.Println()
	root.Reverse()
	root.PreOrder()
	fmt.Println()

	new := root.Copy()
	new.PreOrder()
	fmt.Println()
}