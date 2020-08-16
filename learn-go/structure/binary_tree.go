package main

import (
	"fmt"
	"reflect"
)

/* 二叉树 */

type TreeNode struct {
	Data 	interface{}
	Left 	*TreeNode
	Right 	*TreeNode
}

func (this *TreeNode) PreOrder() {
	if this == nil {
		return
	}
	fmt.Println(this.Data)
	this.Left.PreOrder()
	this.Right.PreOrder()
}

func (this *TreeNode) MidOrder() {
	if this == nil {
		return
	}
	this.Left.MidOrder()
	fmt.Println(this.Data)
	this.Right.MidOrder()
}

func (this *TreeNode) RearOrder() {
	if this == nil {
		return
	}
	this.Left.RearOrder()
	this.Right.RearOrder()
	fmt.Println(this.Data)
}

func (this *TreeNode) Height() int {
	if this == nil {
		return 0
	}

	lh := this.Left.Height()
	rh := this.Right.Height()

	if lh > rh {
		lh++
		return lh
	} else {
		rh++
		return rh
	}
}

func (this *TreeNode) LeafCount(num *int) {
	if this == nil {
		return
	}

	if this.Left == nil && this.Right == nil {
		*num++
	}

	this.Left.LeafCount(num)
	this.Right.LeafCount(num)
}

func (this *TreeNode) Search(data interface{}) {
	if this == nil {
		return
	}

	if reflect.TypeOf(this.Data) == reflect.TypeOf(data) && this.Data == data {
		fmt.Println("数据存在: ", data)
		return
	}

	this.Left.Search(data)
	this.Right.Search(data)
}

func (this *TreeNode) Destroy() {
	if this == nil {
		return
	}

	this.Left.Destroy()
	this.Left = nil

	this.Right.Destroy()
	this.Right = nil

	this.Data = nil
}

func (this *TreeNode) Reverse() {
	if this == nil {
		return
	}

	this.Left, this.Right = this.Right, this.Left

	this.Left.Reverse()
	this.Right.Reverse()
}

func (this *TreeNode) Copy() *TreeNode {
	if this == nil {
		return nil
	}

	left := this.Left.Copy()
	right := this.Right.Copy()

	return &TreeNode{
		Data: this.Data,
		Left: left,
		Right: right,
	}
}

func main() {
	node := &TreeNode{Data: 0}
	node.Left = &TreeNode{Data: 1}
	node.Right = &TreeNode{Data: 2}
	node.Left.Left = &TreeNode{Data: 3}
	node.Left.Right = &TreeNode{Data: 4}
	node.Right.Left = &TreeNode{Data: 5}
	node.Right.Right = &TreeNode{Data: 6}

	node.PreOrder()
	node.MidOrder()
	node.RearOrder()

	height := node.Height()
	fmt.Println(height)

	num := 0
	node.LeafCount(&num)
	fmt.Println(num)

	node.Search(1)

	node.Reverse()
	node.PreOrder()

	newNode := node.Copy()
	newNode.PreOrder()
}
