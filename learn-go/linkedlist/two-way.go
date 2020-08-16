package main

import (
	"fmt"
)

type LinkNode struct {
	Data 	interface{}
	Prev	*LinkNode
	Next	*LinkNode
}

func NewLinkList(datas ...interface{}) *LinkNode {
	if len(datas) == 0 {
		return nil
	}

	head := new(LinkNode)
	point := head

	for _, data := range datas {
		node := new(LinkNode)

		node.Data = data
		node.Prev = point
		node.Next = nil

		point.Next = node
		point = point.Next
	}

	return head
}

func (node *LinkNode) Print() {
	if node == nil {
		return
	}

	for node.Next != nil {
		node = node.Next
		fmt.Println(node.Data)
	}
}

func (node *LinkNode) ReverPrint() {
	if node == nil {
		return
	}

	for node.Next != nil {
		node = node.Next
	}

	for node.Prev != nil {
		fmt.Println(node.Data)
		node = node.Prev
	}
}

func (node *LinkNode) Length() (len int) {
	if node == nil {
		return -1
	}

	for node.Next != nil {
		node = node.Next
		len++
	}

	return
}

func (node *LinkNode) Insert(index int, data interface{}) {
	if node == nil {
		return
	}

	for i := 0; i < index; i++ {
		node = node.Next
	}
	prev := node.Prev
	
	new := new(LinkNode)
	new.Data = data
	new.Next = node
	
	node.Prev = new
	prev.Next = new
	new.Prev = prev
}

func (node *LinkNode) Delete(index int) {
	if node == nil {
		return
	}

	for i := 0; i < index; i++ {
		node = node.Next
	}
	
	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev
}

func (node *LinkNode) Destroy() {
	if node == nil {
		return
	}

	node.Next.Destroy()

	node.Next = nil
	node.Prev = nil
	node.Data = nil
	node = nil
}

func main() {
	head := NewLinkList("java", "python", "golang")
	head.Print()
	// head.ReverPrint()
	fmt.Println(head.Length())

	head.Insert(2, "c/c++")
	head.Print()
	fmt.Println(head.Length())

	head.Delete(3)
	head.Print()
	fmt.Println(head.Length())
}