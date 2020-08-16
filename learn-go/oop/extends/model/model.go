package model

import "fmt"

type Person struct {
	Height int
	Weight int 
}

type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("name=%v, age=%v, score=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupil struct {
	// 嵌入了Student匿名结构体
	// 使用匿名结构体实现OOP编程的继承特性，此时Pupil结构体内部存在Student的属性和方法
	// 结构体可以使用嵌套匿名结构体所有的字段和方法(不论大小写)
	Student
	Person
}

type Graduate struct {
	Student
	// 嵌入有名结构体，这种模式称为组合
	Per Person
}