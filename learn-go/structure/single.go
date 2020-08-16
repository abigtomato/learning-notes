package main

import "fmt"

/* 单链表 */

 type HeroNode struct {
	no 		int			// 节点编号
	name 	string
	next 	*HeroNode	// 指向下一个节点的后继指针
}

// 尾部追加节点
func AppendNode(head *HeroNode, newHeroNode *HeroNode) {
	// 临时指针，用于后移遍历整个链表的节点
	var temp *HeroNode
	temp = head

	for {
		// 若当前指向节点的后继指针指向nil，则代表链表已遍历至末尾
		if (*temp).next == nil {
			break
		}

		// 向后移动指针
		temp = (*temp).next
	}

	// 执行到此则代表已经退出循环，temp指针指向的是最后一个节点，直接将新节点挂在后面即可
	(*temp).next = newHeroNode
}

// 从链表的中间插入节点（按编号排序）
func InsertNode(head *HeroNode, newHeroNode *HeroNode) {
	var temp *HeroNode
	temp = head

	for {
		if (*temp).next == nil {
			break
		} else if (*temp).next.no >= newHeroNode.no {
			// 断开前后两节点的连接，中间插入新节点并重新建立连接
			newHeroNode.next = (*temp).next
			(*temp).next = newHeroNode
			
			return
		}
		
		temp = (*temp).next
	}

	// 执行到此则代表已经退出循环，temp指针指向的是最后一个节点，直接将新节点挂在后面即可
	(*temp).next = newHeroNode
}

// 获取节点
func GetNodeByName(head *HeroNode, name string) (node *HeroNode) {
	var temp *HeroNode
	temp = head

	for {
		if (*temp).next == nil {
			break
		} else if (*temp).next.name == name {
			node = (*temp).next
			return
		}

		temp = (*temp).next
	}

	return
}

// 删除节点
func DelNodeByNo(head *HeroNode, no int) {
	var temp *HeroNode
	temp = head

	for {
		if (*temp).next == nil {
			break
		} else if (*temp).next.no == no {
			// 直接将待删除节点的前一个节点的next指针指向待删除节点的后一个节点即可
			(*temp).next = (*temp).next.next
			return
		}

		temp = (*temp).next
	}
}

// 更新节点信息
func UpdateNode(head *HeroNode, updateHeroNode *HeroNode) {
	var temp *HeroNode
	temp = head

	for {
		if (*temp).next == nil {
			break
		} else if (*temp).next.no == updateHeroNode.no {
			updateHeroNode.next = (*temp).next.next 
			(*temp).next = updateHeroNode
			
			return
		}

		temp = (*temp).next
	}
}

// 展示链表数据
func ShowLinkedList(head *HeroNode) {
	var temp *HeroNode
	temp = head

	if IsEmpty(head) {
		fmt.Printf("IsEmpty(head) fail, linkedlist empty")
		return
	}

	for {
		if (*temp).next == nil {
			break
		}
		fmt.Printf("[val: %v, ptr: %p]--->", (*temp).next.name, (*temp).next.next)
		temp = (*temp).next
	}
}

// 判空
func IsEmpty(head *HeroNode) bool {
	return head.next == nil
}

func main() {
	head := &HeroNode{}

	AppendNode(head, &HeroNode{no: 1, name: "albert",})
	AppendNode(head, &HeroNode{no: 2, name: "lily",})
	AppendNode(head, &HeroNode{no: 3, name: "charname",})

	ShowLinkedList(head)
}