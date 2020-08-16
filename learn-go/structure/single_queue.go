package main

import (
	"fmt"
	"errors"
)

/* 数组实现单向队列 */

 type Queue struct {
	MaxSize 	int				// 最大容量
	Values 		[]interface{}	// 底层数组
	Font 		int				// 头下标指向
	Rear 		int				// 尾下标指向
}

// 初始化一个队列
func NewQueue(maxSize int) (queue *Queue) {
	queue = &Queue{
		MaxSize: maxSize,
		Values: make([]interface{}, maxSize),
		Font: -1,
		Rear: -1,
	}
	return
}

// 入队
func (this *Queue) Add(value interface{}) (err error) {
	// 若尾下标指向了底层数组最后一个元素，则代表队列已满
	if this.Rear == this.MaxSize - 1 {
		err = errors.New("add fail queue full")
		return
	}

	// 尾下标后移，从队尾入队
	this.Rear++
	this.Values[this.Rear] = value

	return
}

// 出队
func (this *Queue) Get() (val interface{}, err error) {
	// 若头下标和尾下标指向相同，则代表队列为空
	if this.Font == this.Rear {
		err = errors.New("get fail queue empty")
		return
	}

	// 头下标后移，从队头出队
	this.Font++
	val = this.Values[this.Font]

	return
}

// 显示队列数据
func (this *Queue) Show() {
	if this.Font == this.Rear {
		return
	}

	// 队列的数据范围从头下标到尾下标
	for i := this.Font; i <= this.Rear; i++ {
		fmt.Printf("queue[%v]=%v\n", i, this.Values[i])
	}
}

func main() {
	queue := NewQueue(10)
	
	for i := 1; i <= 5; i++ {
		queue.Add(i)
	}

	if val, err := queue.Get(); err != nil {
		fmt.Printf("queue.Get() fail error: %v\n", err.Error())
		return
	} else {
		fmt.Printf("val=%v\n", val)
	}

	queue.Show()
}