package main

import (
	"fmt"
	"go_code/learn-go/extends/model"
)

func main() {
	// 操作嵌套匿名结构体的属性方法
	pupil := model.Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.Student.SetScore(30)
	pupil.Student.ShowInfo()

	// 1.编译器会查找结构体中有没有对应的属性和方法，如果没有会继续在嵌套的匿名结构体中查找
	// 2.当结构体和匿名结构体有相同字段和方法时，会将采用就近原则访问，如果希望访问匿名结构体中的数据，则显式调用
	// 3.结构体嵌入多个匿名结构体，如果两个匿名结构体都存在匿名的结构体和方法(同时结构体本身没有)，那么在访问时就必须指定匿名结构体的名字
	graduate := &(model.Graduate{})
	// 4.如果是组合调用，那么必须通过有名结构体名调用
	graduate.Per.Height = 180
	graduate.Name = "albert"
	graduate.Age = 18
	graduate.SetScore(80)
	graduate.ShowInfo()

	// 5.嵌套匿名结构体后，可以在实例时直接指定匿名结构体字段的值
	graduate2 := model.Graduate{
		Student{
			Name: "albert"
			Age: 21
			Score: 100
		}
	}
	fmt.Println(graduate2)
}