package queue

import "fmt"

// cmd> godoc -http :6060 在6060端口开启一个文档服务
// 以"Example + 结构体名 + _ + 方法名"命名的函数被解析为例子
func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
	// 3
	// true
}