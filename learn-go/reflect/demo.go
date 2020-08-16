package main

import (
	"fmt"
	"reflect"
)

/* 
	反射案例
 */

type Monster struct {
	Name string `json:"monster_name"`
	Age int `json:"monster_age"`
	Score float32 `json:"monster_score"`
	Sex string `json:"monster_sex"`
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

// 测试反射操作结构体字段和方法
func TestStruct(a interface{}) {
	rType := reflect.TypeOf(a)
	rValue := reflect.ValueOf(a)
	if kind := rValue.Kind(); kind != reflect.Struct {
		fmt.Println("不是结构体类别")
		return 
	}
	
	// 获取结构体的字段数量
	numOfField := rValue.NumField()
	fmt.Printf("结构体存在%v个字段\n", numOfField)
	

	// 遍历结构体所有字段
	for i := 0; i < numOfField; i++ {
		// 字段索引是按照字段在结构体中定义的顺序决定的
		// rValue.Field(i) 通过索引获取字段的值
		fmt.Printf("字段索引=%v, 字段值=%v\n", i, rValue.Field(i))
		
		// rType.Field(i) 通过索引获取的是StructField结构体类型变量
		// 该变量存在Tag字段，可以直接访问，Get()方法是通过key取value(tag是k-v格式)
		if tagVal := rType.Field(i).Tag.Get("json"); tagVal != "" {
			fmt.Printf("字段索引=%v, 字段标签=%v\n", i, tagVal)
		}
	}

	// 获取结构体的方法数量
	numOfMethod := rValue.NumMethod()
	fmt.Printf("结构体存在%v个方法\n", numOfMethod)

	// rValue.Method(1) 获取结构体的第2个方法(下标从0开始)
	// 方法的索引顺序按照方法名首字母的ASCII码排序
	rValue.Method(1).Call(nil)
	
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	// 通过反射调用方法需要的参数通过 []reflect.Value 切片传递
	// 通过反射调用方法的返回值也是 []reflect.Value 切片类型(兼容多个返回值)
	res := rValue.Method(0).Call(params)
	fmt.Printf("res=%v\n", res[0].Int())
}

// 测试反射修改结构体字段的值
func TestStruct2(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return 
	}

	num := val.Elem().NumField()
	// 获取字段并设置值
	val.Elem().Field(0).SetString("lily")
	for i := 0; i < num; i++ {
		fmt.Printf("索引=%v, 数据=%v, 类型=%v\n",
			i, val.Elem().Field(i), val.Elem().Field(i).Kind())
	}
}

// 编写适配器函数作为函数的统一处理接口
func TestStruct3() {
	call1 := func(v1, v2 int) {
		fmt.Println(v1, v2)
	}

	call2 := func(s1, s2 string) {
		fmt.Println(s1, s2)
	}
	
	// 适配器函数
	bridge := func(call interface{}, args ...interface{}) {
		var params []reflect.Value = make([]reflect.Value, len(args))
		for i := 0; i < len(args); i++ {
			params[i] = reflect.ValueOf(args[i])
		}

		var function reflect.Value = reflect.ValueOf(call)
		function.Call(params)
	}

	bridge(call1, 1, 2)
	bridge(call2, "Hadoop", "Spark")
}

type User struct {
	UserId string
	Name string
}

// 使用反射创建结构体
func TestStruct4(model *User) *User {
	st := reflect.TypeOf(model)
	fmt.Printf("reflect.TypeOf指针类别: %v\n", st.Kind().String())
	fmt.Printf("reflect.TypeOf.Elem结构体类别: %v\n", st.Elem().Kind().String())
	fmt.Println()

	// reflect.New()返回一个Value类型的变量，该值持有一个指向类型为struct的指针
	elem := reflect.New(st.Elem())
	fmt.Printf("reflect.New指针类别: %v\n", elem.Kind().String())
	fmt.Printf("reflect.New.Elem结构体类别: %v\n", elem.Elem().Kind().String())
	fmt.Println()

	// 通过Interface()和类型断言转换成原始类型
	model = elem.Interface().(*User)
	fmt.Printf("原始类型的结构体指针 => 类型=%T, 指向=%p\n", model, model)
	fmt.Println()

	// 取得指针指向的值就是新的结构体实例
	elem = elem.Elem()
	// 为结构体各个字段赋值
	elem.FieldByName("UserId").SetString("123456")
	elem.FieldByName("Name").SetString("Albert")


	return model
}

func main() {
	var a Monster = Monster{
		Name: "albert",
		Age: 400,
		Score: 30.8,
		Sex: "man",
	}

	// 不需要显式的通过结构体类型变量调用方法和操作属性
	// 通过反射机制可以在运行时动态的调用
	TestStruct(a)

	// 通过反射设置结构体字段的值，需要传递指针操作
	TestStruct2(&a)

	// 通过反射机制创建适配器函数
	TestStruct3()

	// 通过反射创建结构体
	var model *User
	model = TestStruct4(model)
	fmt.Printf("创建的新结构体 => model.UserId=%v, model.Name=%v\n", model.UserId, model.Name)
}