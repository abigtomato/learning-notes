package main

import (
	"fmt"
	"errors"
)

 /* 数组实现栈 */

type ArrayStack struct {
	MaxSize		int
	Top 		int
	Values 		[]interface{}
}

func (this *ArrayStack) Push(value interface{}) error {
	if this.Top == this.MaxSize - 1 {
		return errors.New("push fail. stack pull")
	}

	this.Top++
	this.Values[this.Top] = value

	return nil
}

func (this *ArrayStack) Pop() (value interface{}, err error) {
	if this.Top == -1 {
		err = errors.New("pop fail. stack empty")
		return
	}

	value = this.Values[this.Top]
	this.Top--

	return
} 

func (this *ArrayStack) List() error {
	if this.Top == -1 {
		return errors.New("list fail. stack empty")
	}

	for i := this.Top; i >= 0; i-- {
		fmt.Printf("->%v\n", this.Values[i])
	}

	return nil
}

func main() {
	astack := &ArrayStack{
		MaxSize: 5,
		Top: -1,
		Values: make([]interface{}, 5),
	}

	astack.Push(1)
	astack.Push(2)
	astack.Push(3)
	astack.Push(4)
	astack.Push(5)

	astack.List()

	if val, err := astack.Pop(); err != nil {
		fmt.Printf("Pop() fail error: %v\n", err.Error())
	} else {
		fmt.Printf("Pop() success value: %v\n", val)
	}
}
