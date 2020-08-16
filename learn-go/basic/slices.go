package main

import "fmt"

/* 切片 */

// []int的形式是接收数组的切片
func UpdateSlices(arr []int) {
	// 操作切片相当于操作数组本身
	arr[0] = 100
}

// 通过引用已经存在的数组创建切片
func CreateSlice1() {
	/*
		内存分析:
		1. slice是引用类型；
		2. slice在内存中的本质是结构体，有3个成员：指向数组中元素的指针，元素个数，最大容量；
		3. slice中的指针指向被引用的元素段中的第一个元素。
	*/
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	
	/*
		[low:high:max]
		1. low: 起始下标；
		2. high: 结束下标（len = high - low）；
		3. max: 容量（不指定则按照原数组计算，cap = max - low）。
	*/
	slice := arr[0:2:4]
	
	fmt.Printf("数据=%v, 类型=%T, 地址=%p, 第一个元素地址=%p, 第二个元素地址=%p, 元素个数=%v, 最大容量=%v\n",
		arr, arr, &arr, &arr[0], &arr[1], len(arr), cap(arr))
	fmt.Printf("数据=%v, 类型=%T, 指针地址=%p, 指针指向的地址=%p, 第一个元素地址=%p, 第二个元素地址=%p, 元素个数=%v, 最大容量=%v\n",
		slice, slice, &slice, slice, &slice[0], &slice[1], len(slice), cap(slice))
	fmt.Println()
}

// 通过make分配内存创建切片
func CreateSlice2() {
	/*
		内存分析:
		1. 根据指定长度在内存中开辟一段连续的存储空间，存放初值0（相当于在内存中创建了一个数组）；
		2. 切片指针指向这块对外不可见的连续存储空间。
	*/
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 16.245
	slice[3] = .3271
	fmt.Printf("数据=%v, 类型=%T, 元素个数=%v, 最大容量=%v\n", slice, slice, len(slice), cap(slice))
	fmt.Println()
}

// 直接声明切片类型创建切片
func CreateSlice3() {
	var slice []int = []int{1, 3, 5}
	fmt.Printf("数据=%v, 类型=%T, 元素个数=%v, 最大容量=%v\n", slice, slice, len(slice), cap(slice))
	fmt.Println()
}

// 遍历切片
func Traversing(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}
	fmt.Println()

	for i, v := range slice {
		fmt.Printf("[index=%v, value=%v] ", i, v)
	}
	fmt.Println()
}

// 字符串切片
func StrSlice() {
	// 字符串底层是[]byte，也可以进行切片处理
	str := "hello@163.com"
	slice := str[6:]
	fmt.Printf("数据=%v, 类型=%T\n", slice, slice)

	// 修改不可变的字符串方式，转换为[]rune切片(会对中文字符特殊处理)
	arr := []rune(str)
	arr[0] = '数'
	str = string(arr)
	fmt.Printf("修改后的字符串=%v\n", str)
}

// 切片案例，生成斐波那契数列
func Fbn(n int) ([]uint64) {
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i - 1] + fbnSlice[i - 2]
	}

	return fbnSlice
}

// 切片去空
func NoEmpty(data []string) []string {
	i := 0
	for _, v := range data {
		if v != "" {
			data[i] = v
			i++
		}
	}

	return data[:i]
}

// 切片去重
func NoSame(data []string) []string {
	out := data[:1]
	for _, v := range data {
		i := 0
		for ; i < len(out); i++ {
			if v == out[i] {
				break
			}
		}

		if i == len(out) {
			out = append(out, v)
		}
	}

	return out
}

// 删除指定位置元素，并保证元素顺序不变
func remove(data []string, index int) []string {
	copy(data[index:], data[index+1:])
	return data[:len(data)-1]
}

func main() {
	// 1.切片是数组的视图，是指向数组一段区间元素的指针，区间为半开半闭区间[)的形式
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("arr[2:6]=%v, arr[:6]=%v, arr[2:]=%v, arr[:]=%v\n", arr[2:6], arr[:6], arr[2:], arr[:])
	fmt.Println()

	// 2.操作切片相当于操作数组被映射的区间内的元素
	fmt.Println("修改前:", arr[:])
	UpdateSlices(arr[:])
	fmt.Println("修改后:", arr[:])
	fmt.Println()

	// 3.切片的切片还是指向原数组的一段元素
	s1 := arr[:5]
	fmt.Printf("数组的切片=%v, 切片的地址=%p, 切片指向的地址=%p\n", s1, &s1, s1)
	s2 := s1[2:]
	fmt.Printf("切片的切片=%v, 切片的切片的地址=%p, 切片的切片指向的地址=%p\n", s2, &s2, s2)

	/*
		切片append底层分析:
		1. 本质就是对底层数组的扩容；
		2. 按照扩容后的大小创建新的数组（若是根据原有数组创建的切片，会对原数组进行改变）；
		3. 将原来的元素拷贝到新数组中；
		4. 切片重新引用新数组。
	*/
	fmt.Printf("追加前原数组=%v\n", arr)
	fmt.Printf("追加前的切片=%v\n", s2)
	s2 = append(s2, 10)
	fmt.Printf("追加后原数组=%v\n", arr)
	fmt.Printf("追加后的切片=%v\n", s2)

	// 4.切片的拷贝操作
	var slice3 []int = []int{1, 2, 3, 4, 5}
	var slice4 []int = make([]int, 10, 10)
	copy(slice4, slice3)
	fmt.Printf("拷贝后的slice3=%v\n", slice3)
	fmt.Printf("拷贝后的slice4=%v\n", slice4)

	// 5.删除中间的元素
	slice5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 = append(slice5[:3], slice5[4:]...)
	fmt.Println("删除下标为3的元素后: ", slice5)

	// 6.删除头部元素
	front := slice5[0]
	slice5 = slice5[1:]
	fmt.Println("头部元素: ", front)
	fmt.Println("删除头部元素后: ", slice5)

	// 7.删除尾部元素
	tail := slice5[len(slice5) - 1]
	slice5 = slice5[:len(slice5) - 1]
	fmt.Println("尾部元素: ", tail)
	fmt.Println("删除尾部元素后: ", slice5)

	// 8.创建切片
	CreateSlice1()
	CreateSlice2()
	CreateSlice3()

	// 9.遍历切片
	Traversing(arr[:])

	// 10.字符串切片
	StrSlice()

	// 11.斐波那契案例
	Fbn(10)	

	// 12.切片去空
	data := []string{"red", "", "black", "", "", "pink", "blue"}
	data = NoEmpty(data)
	fmt.Printf("%v, %d, %d\n", data, len(data), cap(data))

	// 13.切片去重
	data = []string{"red", "black", "red", "red", "blue", "pink", "red"}
	data = NoSame(data)
	fmt.Printf("%v, %d, %d\n", data, len(data), cap(data))

	// 14.指定位置删除，并保证原顺序不变
	data = []string{"hadoop", "hive", "hbase", "docker", "spark", "flink", "kafka"}
	data = remove(data, 2)
	fmt.Printf("%v, %d, %d\n", data, len(data), cap(data))
}