package main

import (
	"fmt"
	"reflect"
)

/* 
	反射基础
 */

// 基本类型反射
func reflectTest(basic interface{}) {
	// 动态获取任意变量的类型信息，返回实现了Type接口的变量(Type接口定义了大量操作变量类型的方法)
	rType := reflect.TypeOf(basic)
	fmt.Printf("数据=%v, 类型=%T, 指向=%p\n", rType, rType, rType)

	// 动态获取任意变量的值信息，返回Value结构体类型的变量(提供了大量用于操作变量值的方法)
	rValue := reflect.ValueOf(basic)
	// 使用反射变量操作原始变量，Int()表示获取原始int类型数据的值 
	value := rValue.Int()
	fmt.Printf("数据=%v, 类型=%T, 原值=%v\n", rValue, rValue, value)

	// 通过Interface()方法返回空接口
	iValue := rValue.Interface()
	// 通过类型断言转换成原始数据类型
	num := iValue.(int)
	fmt.Printf("数据=%v, 类型=%T\n", num, num)
}

type Student struct {
	Name string
	Age int
}

// 结构体反射
func reflectTest2(basic interface{}) {
	rType := reflect.TypeOf(basic)
	fmt.Printf("数据=%v, 类型=%T, 指向=%p\n", rType, rType, rType)

	rValue := reflect.ValueOf(basic)
	fmt.Printf("数据=%v, 类型=%T\n", rValue, rValue)

	// Kind()方法获取变量所属的类别(如果是结构体，那么类别就是struct)
	kind1 := rType.Kind()
	kind2 := rValue.Kind()
	fmt.Printf("Type获取类别=%v, Value获取类别=%v\n", kind1, kind2)

	iValue := rValue.Interface()
	if stu, ok := iValue.(Student); ok {
		fmt.Printf("数据=%v, 类型=%T, 字段Name=%v, 字段Age=%v\n", stu, stu, stu.Name, stu.Age)
	}
}

// 通过反射改变基本类型的值
func reflectTest3(basic interface{}) {
	rVal := reflect.ValueOf(basic)
	fmt.Printf("类别=%v, 类型=%T\n", rVal.Kind(), rVal)

	// rVal是指针，Elem()是取指针指向的数据
	rVal.Elem().SetInt(20)
	fmt.Printf("rVal.Elem()->val=%v, rVal.Elem()->type=%T\n", rVal.Elem(), rVal.Elem())
}

func main() {
	var num int = 10
	reflectTest(num)
	fmt.Println()

	var stu = Student{
		Name: "Tom",
		Age: 20,
	}
	reflectTest2(stu)
	fmt.Println()

	// 想要通过反射修改值，需要指针类型
	var pNum *int = &num
	reflectTest3(pNum)
	fmt.Printf("通过反射改变后=%v\n", num)
}