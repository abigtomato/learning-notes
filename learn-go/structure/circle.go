package main

import "fmt"

/* 循环链表 */

type CatNode struct {
	no 		int
	name 	string
	next 	*CatNode
}

// 插入节点
func InsterCatNode(head *CatNode, newNode *CatNode) {
	// 头结点的情况(head节点拥有数据)
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		// head节点指向自己(单节点循环)
		head.next = head
		return
	}

	temp := head
	for {
		// 若是当前节点的下一个节点是head节点，那么该节点是最后一个节点
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	// 从尾部插入，新插入的节点next指针指向head节点形成环形
	temp.next = newNode
	newNode.next = head
}

// 展示环形链表的数据
func CircleLinkedList(head *CatNode) {
	temp := head

	if temp.next == nil {
		return
	}

	for {
		fmt.Printf("[no=%v, name=%v, next=%p]--->", temp.no, temp.name, temp.next)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

// 删除环形链表的节点
func DelCatNode(head *CatNode, name string) *CatNode {
	// 1.temp指针指向head节点，用于循环比较各节点值和待删除节点值是否相同
	// 2.helper指针指向最后一个节点，随temp的移动而移动，永远指向temp指向的节点的前一个节点(协助删除节点)
	temp, helper := head, head

	// 循环链表为空
	if temp.next == nil {
		return head
	}

	// 单节点的循环链表
	if temp.next == head {
		temp.next = nil
		return head
	}

	// 初始化helper指针的指向(指向最后一个节点)
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	for {
		// 最后一个节点的情况
		if temp.next == head {
			// 比较最后一个节点
			if temp.name == name {
				helper.next = temp.next
			} else {
				fmt.Println("没有找到待删除的节点")
			}
			break
		} else if temp.name == name {
			// 若待删除节点是head节点
			if temp == head {
				// 后移head节点，并将新的头节点返回
				// 因为main栈中还存在一个head指针指向head节点，这里只是改变了该方法栈的head指针，需要返回值覆盖
				head = head.next
			} 
			// 通过helper指针协助temp指针删除节点
			helper.next = temp.next
			break
		}
		// helper指针随temp指针移动而移动，永远指向temp指向的节点的前一个节点
		temp = temp.next
		helper = helper.next
	}

	return head
}

func main() {
	head := &CatNode{}
	InsterCatNode(head, &CatNode{no: 1, name: "spark core",})
	InsterCatNode(head, &CatNode{no: 2, name: "spark sql",})
	InsterCatNode(head, &CatNode{no: 3, name: "spark streaming",})
	InsterCatNode(head, &CatNode{no: 4, name: "spark mllib",})
	InsterCatNode(head, &CatNode{no: 5, name: "spark graphx",})
	
	CircleLinkedList(head)
	fmt.Println()

	head = DelCatNode(head, "spark streaming")
	CircleLinkedList(head)
	fmt.Println()
}