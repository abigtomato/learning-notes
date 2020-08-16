package main

import (
	"fmt"
	"reflect"
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

	for i := 0; i < len(datas); i++ {
		node := &LinkNode{
			Data: datas[i],
			Next: nil,
		}

		point.Next = node
		point = point.Next
	}

	return head
}

func (node *LinkNode) Printf() {
	if node.Next == nil {
		return
	}

	for node.Next != nil {
		node = node.Next
		fmt.Printf("[%v]\n", node.Data)
	}
}

func (node *LinkNode) Println() {
	if node == nil {
		return
	}

	fmt.Println(node.Data)
	
	node.Next.Println()
}

func (node *LinkNode) Length() (len int) {
	if node == nil {
		len = -1
		return
	}

	for node.Next != nil {
		node = node.Next
		len++
	}

	return
}

func (node *LinkNode) InsertByHead(data interface{}) {
	if node == nil {
		return
	}

	if data == nil {
		return
	}

	new := new(LinkNode)
	new.Data = data
	new.Next = node.Next
	
	node.Next = new
}

func (node *LinkNode) InsertByTail(data interface{}) {
	if node == nil {
		return 
	}

	if data == nil {
		return
	}

	new := new(LinkNode)
	new.Data = data
	new.Next = nil

	for node.Next != nil {
		node = node.Next
	}

	node.Next =  new
}

func (node *LinkNode) InsertByIndex(index int, data interface{}) {
	if node == nil {
		return 
	}

	if index < 0 || index > node.Length() {
		return
	}

	if data == nil {
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

func (node *LinkNode) DeleteByIndex(index int) {
	if node == nil {
		return
	}

	if index < 0 || index > node.Length() {
		return
	}

	for i := 0; i < index - 1; i++ {
		node = node.Next
	}

	node.Next = node.Next.Next
}

func (node *LinkNode) DeleteByData(data interface{}) {
	if node == nil {
		return
	}

	if data == nil {
		return
	}

	pre := node
	for node.Next != nil {
		pre = node
		node = node.Next
		
		if reflect.TypeOf(node.Data) == reflect.TypeOf(data) && node.Data == data {
			pre.Next = node.Next
			return		
		}
	}
}

func (node *LinkNode) GetIndexByData(data interface{}) (index int) {
	if node == nil {
		return -1
	}

	if data == nil {
		return -1
	}

	for node.Next != nil {
		node = node.Next
		index++
		
		if reflect.TypeOf(node.Data) == reflect.TypeOf(data) && reflect.DeepEqual(node.Data, data) {
			return
		}
	}

	return -1
}

func (node *LinkNode) GetDataByIndex(index int) (data interface{}) {
	if node == nil {
		return
	}

	if index < 0 || index > node.Length() {
		return
	}

	for i := 0; i < index; i++ {
		node = node.Next
	}	

	return node.Data
}

func (node *LinkNode) Destroy() {
	if node == nil {
		return
	}

	node.Next.Destroy()

	node.Data = nil
	node.Next = nil
	node = nil
}

func main() {
	head := NewLinkList("hadoop", "spark", "flink")
	head.Printf()
	fmt.Println(head.Length())

	head.InsertByHead("hive")
	head.Printf()
	fmt.Println(head.Length())

	head.InsertByTail("hbase")
	head.Printf()
	fmt.Println(head.Length())

	head.InsertByIndex(3, "zookeeper")
	head.Printf()
	fmt.Println(head.Length())

	head.DeleteByIndex(3)
	head.Printf()
	fmt.Println(head.Length())

	head.DeleteByData("spark")
	head.Printf()
	fmt.Println(head.Length())

	index := head.GetIndexByData("flink")
	fmt.Println(index)

	head.Printf()
	data := head.GetDataByIndex(3)
	fmt.Println(data)
}