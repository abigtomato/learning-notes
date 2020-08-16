package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/* 函数 */

// 全局变量，首字母大写表示在整个程序有效
var Name string
// 包变量在整个包中有效
var Age int
// 全局匿名函数
var fun = func(a, b int) (c, d int) {
	return a / b, a % b 
}

// 初始化函数，一个go文件的执行顺序是定义全局变量 -> 执行init()函数 -> 执行main()函数
func init() {
	Name = "Albert"
	Age = 21
	fmt.Println(Name, Age)
	
	ret1, ret2 := fun(9, 4)
	fmt.Println(ret1, ret2)
}

// 第一个()表示参数名和类型，第二个()表示返回值类型，可存在多个返回值
func Eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("Unsupported Operation: %s", op)
	}
}

// 定义返回值的名称
func Div(a, b int) (q, r int) {
	// 执行到defer时，会将修饰的语句和相关的值压入defer栈中
	// 当函数执行完毕后，从defer栈顶弹出语句执行
	defer fmt.Println("先压入，后弹出")
	defer fmt.Println("后压入，先弹出")

	q = a / b
	r = a % b

	return q, r
}

// 闭包，函数做为参数传递
func Apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("函数名:%s, 参数:(%d, %d)\n", opName, a, b)

	return op(a, b)
}

// ...可变参数传递
func SumArgs(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	
	// 注: 函数不能返回局部变量的地址值；
	// 局部变量保存在栈帧上，函数调用结束后释放栈帧，局部变量的地址不再受系统保护，随时可能分配给其他程序。
	return sum
}

// 引用传递，*类型表示接收对应类型变量的内存地址
func Swap(a, b *int) {
	// *int 星号+类型表示指向int类型变量的指针(那么a和b则是指针类型的变量)
	// *a 星号+指针类型变量表示取指针指向的变量值
	fmt.Println("指针:", a)
	fmt.Println("*指针:", *a)

	/*
		等式分析:
		1. 等式左边的值表示a，b地址指向的内存空间；
		2. 等式右边的值表示a，b地址指向空间中的数据；
		3. 将a，b指向空间中的数据存入b，a指向的空间中。
	*/
	*a, *b = *b, *a	// 使用指针交互两个变量的值
}

// 递归函数解决斐波那契数列
func Fib(total int) int {
	if total == 1 || total == 2 {
		return 1
	} else {
		return fib(total - 1) + fib(total - 2)
	}
}

func Peach(total int) int {
	if total == 10 {
		return 1
	}

	return (peach(total + 1) + 1) * 2
}

func main() {
	/*
		函数调用机制底层分析: 
		1. 基本数据类型一般分配到栈区，编译器存在逃逸分析，会自动判断什么时候分配到什么内存区域；
		2. 引用数据类型一般分配到堆区。
		
		代码存放在内存中的代码区:
		1. 在调用函数时，会在栈区为当前函数分配一段存储空间，编辑器会自动处理让此内存空间和其他栈空间区分开来；
		2. 每个函数对应的栈空间是相对独立的；
		3. 若是值传递，则会拷贝参数存放在函数独立的内存区域，若是引用传递，在函数独立空间只会开辟指针变量接收原数据的地址；
		4. 程序先执行main函数，将main函数独立的空间压入栈空间，当其他函数被调用，会分配空间并压入栈，函数执行结束会从栈顶释放对应的空间。
	*/
	if res, err := Eval(3, 4, "*"); err == nil {
		fmt.Println(res)
	}

	// 1.函数是一种数据类型，可以赋值给变量，该变量就是函数类型的变量
	fun := div
	fmt.Printf("fun的类型:%T, div的类型:%T\n", fun, div)
	q, r := fun(3, 4)
	fmt.Println(q, r)

	// 2.传递匿名函数做为参数
	ret := Apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4)
	fmt.Println(ret)

	// 3.可变参数传递
	sum := SumArgs(1, 2, 3, 4, 5)
	fmt.Println(sum)

	// 4.引用传递
	a, b := 3, 4
	Swap(&a, &b)	// &为取地址符，&a和&b表示传递a，b变量的内存地址
	fmt.Println(a, b)

	// 5.递归函数
	ret = Fib(5)
	fmt.Println(ret)

	// 6.
	ret = Peach(5)
	fmt.Println(ret)

	// 7.new()函数用于分配值类型的内存
	num := new(int)
	fmt.Printf("值=%v, 类型=%T, 地址=%v, 指向的值=%v\n", num, num, &num, *num)
}