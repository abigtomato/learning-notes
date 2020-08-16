package main

import "fmt"

/* 环形链表解决约瑟夫环问题 */

type BoyNode struct {
	no 		int
	next 	*BoyNode
}

// 生成环形链表
func CreateCircle(num int) *BoyNode {
	// first指针始终指向head节点
	// curBoy指针向后移动遍历节点元素
	var first, curBoy *BoyNode
	
	if num <= 0 {
		return nil
	}

	for i := 1; i <= num; i++ {
		boy := &BoyNode{no: i,}
		// 第一个节点的情况
		if i == 1 {
			// 两个指针初始化都指向第一个节点
			first = boy
			curBoy = boy
			// 指向自己形成环形
			curBoy.next = first
		} else {
			// 非第一个节点的情况，在最后追加
			curBoy.next = boy
			curBoy = boy
			boy.next = first
		}
	}
	
	return first
}

// 展示环形数据
func ShowCircle(first *BoyNode) {
	var curBoy *BoyNode
	curBoy = first

	for {
		fmt.Printf("[no: %v, next: %p]--->", curBoy.no, curBoy.next)
		if curBoy.next == first {
			break
		}
		curBoy = curBoy.next
	}
}

// 开始约瑟夫环游戏
func PlayGame(first *BoyNode, startNo int, coutNum int) {
	// 辅助删除的节点
	var tail *BoyNode = first

	if first.next == nil {
		fmt.Println("circle empty")
		return
	}

	// 初始化辅助节点(指向链表最后一个节点)
	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}
	
	// 移动两个指针指向游戏开始的位置
	// 因为当前节点也计数，所有移动步数减一
	for i := 1; i <= startNo - 1; i++ {
		first = first.next
		tail = tail.next
	}

	for {
		// 按照规则，tail指针每次移动一定的步数，都会移除当前指向的节点
		for i := 1; i <= coutNum - 1; i ++ {
			first = first.next
			tail = tail.next
		}
		
		fmt.Printf("第%v号boy出列\n", first.no)
		// 通过tail指针的辅助删除被选中的节点
		first = first.next
		tail.next = first

		// 若tail指针和first指针移动到同一个位置，说明链表当前只有一个节点
		if tail == first {	
			break
		}
	}
	// 出队最后的一个节点
	fmt.Printf("第%v号boy出列\n", first.no)
}

func main() {
	first := CreateCircle(10)
	
	ShowCircle(first)
	fmt.Println()
	
	PlayGame(first, 2, 3)
}