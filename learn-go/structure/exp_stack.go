package main

import (
	"fmt"
	"errors"
	"strconv"
)

/* 使用栈进行表达式计算 */

type Stack struct {
	MaxSize 	int
	Top 		int
	Values		[]interface{}
}

func (this *Stack) Push(value interface{}) error {
	if this.Top == this.MaxSize - 1 {
		return errors.New("push fail. stack pull")
	}

	this.Top++
	this.Values[this.Top] = value

	return nil
}

func (this *Stack) Pop() (value interface{}, err error) {
	if this.Top == -1 {
		err = errors.New("pop fail. stack empty")
		return
	}

	value = this.Values[this.Top]
	this.Top--

	return
} 

func (this *Stack) List() error {
	if this.Top == -1 {
		return errors.New("list fail. stack empty")
	}

	for i := this.Top; i >= 0; i-- {
		fmt.Printf("->%v\n", this.Values[i])
	}

	return nil
}

func IsOper(val int) bool {	
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

func IsNum(val int) bool {
	if val >= 48 && val <= 57 {
		return true
	} else {
		return false
	}
}

func Cal(num1 int, num2 int, oper int) (res int, err error) {
	switch oper {
		case 42:
			res = num2 * num1
		case 43:
			res = num2 + num1
		case 45:
			res = num2 - num1
		case 47:
			res = num2 / num1
		default:
			err = errors.New("oper error")
	}
	return
}

func Priority(oper int) (res int) {
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return
}

func Calculation(numStack, operStack *Stack, exp string) (res int, err error) {
	chExp := []byte(exp)
	var keepNum string
	
	for i, ch := range chExp {
		if IsOper(int(ch)) {
			if operStack.Top == -1 {
				operStack.Push(int(ch))
			} else {
				if Priority(int(ch)) >= Priority(operStack.Values[operStack.Top].(int)) {
					operStack.Push(int(ch))
				} else {
					num1, _ := numStack.Pop()
					num2, _ := numStack.Pop()
					oper, _ := operStack.Pop()
					res, err = Cal(num1.(int), num2.(int), oper.(int))
					numStack.Push(res)
					operStack.Push(int(ch))
				}
			}
		} else if IsNum(int(ch)) {
			keepNum += fmt.Sprintf("%c", ch)
			if i == len(chExp) - 1 {
				num, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(num))
			} else {
				if IsOper(int(chExp[i + 1])) {
					num, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(num))
					keepNum = ""
				}
			}
		}
	}

	for operStack.Top != -1 {
		num1, _ := numStack.Pop()
		num2, _ := numStack.Pop()
		oper, _ := operStack.Pop()
		res, err = Cal(num1.(int), num2.(int), oper.(int))
		numStack.Push(res)
	}

	num, err := numStack.Pop()
	res = num.(int)

	return
}

func main() {
	numStack := &Stack{
		MaxSize: 20,
		Top: -1,
		Values: make([]interface{}, 20),
 	}
	operStack := &Stack{
		MaxSize: 20,
		Top: -1,
		Values: make([]interface{}, 20),
	}
	
	res, err := Calculation(numStack, operStack, "30+20*60-200")
	if err != nil {
		fmt.Printf("Calculation() fail error: %v\n", err.Error())
		return
	}
	fmt.Printf("result: %v\n", res)
}