package main

import "fmt"

 /* 双向链表 */

type Node struct {
	no 		int
	name 	string
	pre 	*Node	// 前趋指针
	next 	*Node	// 后继指针
}

type DoubleLinkedList struct {
	head *Node
}

// 链表末尾追加节点
func (this *DoubleLinkedList) AppendNode(newNode *Node) {
	// 临时指针，用于遍历节点
	var temp *Node
	temp = this.head

	for {
		if (*temp).next == nil {
			break
		}

		temp = (*temp).next
	}

	// 此时temp指向最后一个节点，使temp指向节点后继指向新节点，新节点前序指向temp指向节点
	(*temp).next = newNode
	newNode.pre = temp
}

// 插入节点
func (this *DoubleLinkedList) InsertNode(newNode *Node) {
	var temp *Node
	temp = this.head

	for {
		if (*temp).next == nil {
			break
		} else if (*temp).next.no >= newNode.no {
			// 断开原节点的连接，插入新节点，重新建立连接
			newNode.pre = (*temp).pre
			(*temp).pre.next = newNode

			(*temp).pre = newNode
			newNode.next = temp

			return
		}

		temp = (*temp).next
	}

	// 节点插入末尾的情况
	(*temp).next = newNode
	newNode.pre = temp
}

// 删除节点
func (this *DoubleLinkedList) DelNodeByName(name string) {
	var temp *Node
	temp = this.head

	for {
		if (*temp).next == nil {
			if (*temp).name == name {
				(*temp).pre.next = nil
			}
			break
		} else if (*temp).name == name {
			(*temp).pre.next = (*temp).next
			(*temp).next.pre = (*temp).next
			
			return
		}

		temp = (*temp).next
	}
}

// 展示链表所有节点
func (this *DoubleLinkedList) ShowLinkedList() {
	var temp *Node
	temp = this.head

	for {
		if (*temp).next == nil {
			break
		}

		fmt.Printf("[no=%v, name=%v, pre=%p, next=%p]-->", (*temp).next.no, (*temp).next.name, (*temp).next.pre, (*temp).next.next)
		temp = (*temp).next
	}
}

// 判空
func (this *DoubleLinkedList) IsEmpty() bool {
	return this.head.next == nil
}

func main() {
	link := &DoubleLinkedList{
		head: &Node{},
	}

	link.AppendNode(&Node{no: 1, name: "hadoop"})
	link.AppendNode(&Node{no: 2, name: "spark",})
	link.AppendNode(&Node{no: 3, name: "kafka",})
	link.AppendNode(&Node{no: 5, name: "storm",})
	link.ShowLinkedList()
	fmt.Println()

	link.InsertNode(&Node{no: 4, name: "hbase",})
	link.InsertNode(&Node{no: 6, name: "hive",})
	link.ShowLinkedList()
	fmt.Println()

	link.DelNodeByName("hive")
	link.DelNodeByName("kafka")
	link.ShowLinkedList()
	fmt.Println()
}