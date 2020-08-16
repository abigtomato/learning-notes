package main

import (
	"fmt"
)

type LinkNode struct {
	Data	interface{}
	Next 	*LinkNode
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
		node.Next = head.Next

		point.Next = node
		point = point.Next
	}

	return head
}

func (node *LinkNode) Print() {
	if node == nil {
		return
	}

	start := node.Next

	for {
		node = node.Next
		fmt.Printf("%v ", node.Data)

		if node.Next == start {
			break
		}
	}
}	

func (node *LinkNode) Length() (len int) {
	if node == nil {
		return -1
	}

	start := node.Next

	for {
		node = node.Next
		len++

		if node.Next == start {
			return
		}
	}
}

func (node *LinkNode) Insert(index int, data interface{}) {
	if node == nil {
		return
	}

	if index < 0 || index > node.Length() + 1 {
		return
	}

	if data == nil {
		return
	}

	// 头插法
	if index == 1 {
		node.insertByHead(data)	
		return
	}

	// 尾插法
	if index == node.Length() + 1 {
		node.insertByTail(data)
		return
	}

	for i := 0; i < index - 1; i++ {
		node = node.Next	
	}

	new := new(LinkNode)
	new.Data = data
	new.Next = node.Next

	node.Next = new
}

func (node *LinkNode) insertByHead(data interface{}) {
	head := node
	start := node.Next
	
	for {
		node = node.Next
		if node.Next == start {
			break
		}
	}

	new := new(LinkNode)
	new.Data = data
	new.Next = head.Next

	head.Next = new
	node.Next = new
}

func (node *LinkNode) insertByTail(data interface{}) {
	start := node.Next
	for {
		node = node.Next
		if node.Next == start {
			break
		}
	}

	new := new(LinkNode)
	new.Data = data
	new.Next = start

	node.Next = new
}

func (node *LinkNode) Delete(index int) {
	if node == nil {
		return
	}	

	if index < 0 || index > node.Length() {
		return
	}

	// 删除头
	if index == 1 {
		head := node
		start := node.Next
		
		for {
			node = node.Next
			if node.Next == start {
				break
			}
		}

		head.Next = start.Next
		node.Next = head.Next

		return
	}

	for i := 0; i < index - 1; i++ {
		node = node.Next
	}

	node.Next = node.Next.Next
}

func Josephu() {
	var datas []interface{}
	for i := 1; i <= 32; i++ {
		datas = append(datas, i)
	}

	head := NewLinkList(datas...)
	head.Print()

	i := 0
	for head.Length() > 2 {
		i += 3
		if i > head.Length() {
			i = head.Length() % 3
		}

		head.Delete(i)
		head.Print()
		fmt.Println()
	}
}

func main() {
	head := NewLinkList("albert", "lily", "charname", "kristen", "king", "queue")
	head.Print()
	fmt.Println(head.Length())

	head.Insert(3, "233")
	head.Print()
	fmt.Println(head.Length())

	head.Delete(2)
	head.Print()
	fmt.Println(head.Length())

	Josephu()
}