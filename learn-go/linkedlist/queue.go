package main

import "fmt"

type QueueNode struct {
	Data 	interface{}
	Next 	*QueueNode
}

func NewQueue(datas ...interface{}) *QueueNode {
	head := new(QueueNode)
	point := head

	for _, data := range datas {
		node := new(QueueNode)
		node.Data = data
		node.Next = point.Next

		point.Next = node
		point = point.Next
	}

	return head
}

func (node *QueueNode) Print() {
	if node == nil {
		return
	}

	for node.Next != nil {
		node = node.Next
		fmt.Println(node.Data)
	}
}

func (node *QueueNode) Length() (len int) {
	if node == nil {
		return -1
	}

	for node.Next != nil {
		node = node.Next
		len++
	}
	
	return
}	

func (node *QueueNode) Enqueue(data interface{}) {
	if node == nil {
		return
	}

	for node.Next != nil {
		node = node.Next
	}
	
	new := new(QueueNode)
	new.Data = data
	new.Next = node.Next

	node.Next = new
}

func (node *QueueNode) Dequeue() interface{} {
	if node == nil {
		return nil
	}

	de := node.Next
	node.Next = node.Next.Next

	return de.Data
}

func main() {
	head := NewQueue("java", "golang", "python")
	head.Print()
	fmt.Println(head.Length())

	head.Enqueue("c/c++")
	head.Print()	
	fmt.Println(head.Length())

	fmt.Println(head.Dequeue())
	fmt.Println(head.Length())
	head.Print()
}