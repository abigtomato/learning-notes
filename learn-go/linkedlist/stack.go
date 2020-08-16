package main

import (
	"fmt"
)

type StackNode struct {
	Data 	interface{}
	Next 	*StackNode
}

func NewStack(datas ...interface{}) *StackNode {
	head := new(StackNode)

	for _, data := range datas {
		node := new(StackNode)
		node.Data = data
		node.Next = head.Next
		
		head.Next = node
	}

	return head
}

func (node *StackNode) Print() {
	if node == nil {
		return
	}

	for node.Next != nil {
		node = node.Next
		fmt.Println(node.Data)
	}
}

func (node *StackNode) Length() (len int) {
	if node == nil {
		return -1
	}

	for node.Next != nil {
		node = node.Next
		len++
	}

	return
}

func (node *StackNode) Push(data interface{}) {
	if node == nil {
		return
	}

	new := new(StackNode)
	new.Data = data
	new.Next = node.Next

	node.Next = new
}

func (node *StackNode) Pop() interface{} {
	if node == nil {
		return nil
	}
	
	pop := node.Next
	node.Next = node.Next.Next

	return pop.Data
}

func main() {
	head := NewStack("java", "python", "golang")
	head.Print()
	fmt.Println(head.Length())

	head.Push("c/c++")
	head.Print()
	fmt.Println(head.Length())

	fmt.Println(head.Pop())
	fmt.Println(head.Length())
	head.Print()
}