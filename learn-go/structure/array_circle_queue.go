package main

import (
	"fmt"
	"errors"
)

/* 数组实现循环队列 */

type CircleQueue struct {
	MaxSize int	
	Values []interface{}
	Head int
	Tail int
}

// 初始化循环队列
func NewCircleQueue(maxSize int) *CircleQueue {
	return &CircleQueue{
		MaxSize: maxSize,
		Values: make([]interface{}, maxSize),
		Head: 0,
		Tail: 0,
	}
}

// 入队
func (this *CircleQueue) Push(val interface{}) (err error) {
	if this.IsFull() {
		err = errors.New("push fail queue full")
		return
	}

	// 下标每次移动后都要与队列最大长度取模
	// (这样才会达成循环的目的，移动到尾部后和maxsize取模会再次指向首部)
	this.Values[this.Tail] = val
	this.Tail = (this.Tail + 1) % this.MaxSize

	return
}

// 出队
func (this *CircleQueue) Pop() (val interface{}, err error) {
	if this.IsEmpty() {
		err = errors.New("pop fail queue empty")
		return
	}

	val = this.Values[this.Head]
	this.Head = (this.Head + 1) % this.MaxSize

	return
}

// 计算队列元素数量
func (this *CircleQueue) Show() {
	size := this.Size()
	
	if size == 0 {
		fmt.Println("queue size = 0")
		return
	}

	// 临时的指向，从头下标开始不断后移遍历所有元素
	tempHead := this.Head
	for i := 0; i < size; i++ {
		fmt.Printf("queue[%v]=%v\n", tempHead, this.Values[tempHead])
		tempHead = (tempHead + 1) % this.MaxSize
	}
}

// 判断队列是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.Head == this.Tail
}

// 判断队列是否已满
func (this *CircleQueue) IsFull() bool {
	return (this.Tail + 1) % this.MaxSize == this.Head
}

// 队列元素个数
func (this *CircleQueue) Size() int {
	return (this.Tail + this.MaxSize - this.Head) % this.MaxSize
}

func main() {
	queue := NewCircleQueue(10)

	for i := 1; i <= 10; i++ {
		queue.Push(i)
	}
	queue.Show()

	for i := 0; i <= 5; i++ {
		val, err := queue.Pop()
		if err != nil {
			break
		}
		fmt.Printf("pop val=%v\n", val)
	}

	for i := 1; i <= 5; i++ {
		queue.Push(i)
	}
	queue.Show()
}