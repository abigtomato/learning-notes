package main

import (
	"fmt"
	"unsafe"
	"encoding/json"
)

type Cat struct {
	Name string	`json:"name"`	// struct tag
	Age int `json:"age"`
	Color string `json:"color"`
	Hobby string `json:"hobby"`
}

func main() {
	// 创建结构体变量
	// 结构体变量的内存分析:
	// 1.结构体变量是值类型
	// 2.在内存中分配一段存储空间，保存结构体内部的各个字段(所有字段在内存中是连续分配的)
	// 3.因为是值类型，所以该段内存中的数据可通过变量名直接访问
	// 4.如果字段是指针类型，指针本身的地址还是连续的，但指向的地址不一定连续
	var cat1 Cat = Cat{
		Name: "小白",
		Age: 3,
		Color: "白色",
		Hobby: "吃<~)))><<",
	}
	fmt.Printf("数据=%v, 类型=%T, 地址=%p, 大小: %v\n", cat1, cat1, &cat1, unsafe.Sizeof(cat1))
	fmt.Println()

	// 创建结构体指针
	var cat2 *Cat = &Cat{
		Name: "肥橘",
		Age: 5,
		Color: "橘色",
		Hobby: "吃<~)))><<",
	}
	fmt.Printf("结构体指针 => 数据=%v, 类型=%T, 指向=%p\n", cat2, cat2, cat2)
	fmt.Printf("结构体变量 => 数据=%v, 类型=%T, 地址=%p\n", *cat2, *cat2, &(*cat2))
	fmt.Println()

	// new()分配内存空间并返回结构体指针
	var cat3 *Cat = new(Cat)
	// 只有通过结构体变量才能通过.操作结构体内属性
	// 使用简化写法cat2.Name = "小黑"也可以
	(*cat3).Name = "小黑"
	(*cat3).Age = 4
	(*cat3).Color = "黑色"
	(*cat3).Hobby = "吃<~)))><<"
	fmt.Printf("结构体指针 => 数据=%v, 类型=%T, 指向=%p\n", cat3, cat3, cat3)
	fmt.Printf("结构体变量 => 数据=%v, 类型=%T, 地址=%p\n", *cat3, *cat3, &(*cat3))
	fmt.Println()

	// 将cat4序列化为json字符串，会根据struct tag来设置json的字段名
	cat4 := Cat{"小花", 5, "黑白相间", "吃<~)))><<"}
	jsonStr, _ := json.Marshal(cat4)
	fmt.Printf("json=%v\n", string(jsonStr))

	//
	var p1 Cat = Cat{"肥橘", 5, "橘色", "吃<~)))><<"}
	var p2 Cat = Cat{"肥橘", 5, "橘色", "吃<~)))><<"}
	var p3 Cat = Cat{"肥橘", 6, "橘色", "吃<~)))><<"}
	fmt.Println("p1 == p2 ? ", p1 == p2)
	fmt.Println("p1 == p3 ? ", p1 == p3)
}