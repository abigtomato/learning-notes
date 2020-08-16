package main

import (
	"fmt"
	"sort"
	"strings"
)

/* 映射 */

// 1.创建map
func Create() map[string]string {
	// 方式一: 创建时赋值（key无序且不能重复，重复会覆盖）
	m := map[string]string {
		"name":		"ccmouse",
		"course": 	"golang",
		"site": 	"imooc",
		"quality": 	"notbad",
	}

	// 方式二: 通过make()分配空间创建map
	m2 := make(map[string]int, 10)

	// 方式三: 声明式创建是不会分配内存空间的
	var m3 map[string]int

	fmt.Printf("m=%v, m2=%v, m3=%v\n", m, m2, m3)
	fmt.Printf("数据: %v, 类型: %T, 地址: %p, 长度: %v\n", m2, m2, m2, len(m2))

	return m
}

// 2.map的增删改查
func CRUD(m map[string]string) {
	// 新增key-value
	m["age"] = "18"	
	fmt.Printf("m=%v\n", m)

	// 新增的key相同，则为更新
	m["age"] = "20"
	fmt.Printf("m=%v\n", m)

	// 按key删除
	delete(m, "name")
	fmt.Printf("m=%v\n", m)

	// 按key取value，第二个返回值表示是否存在value
	if site, ok := m["quality"]; !ok {
		fmt.Println("key does not exist ......")
	} else {
		fmt.Printf("site=%v\n", site)
	}

	// 清空key
	m = make(map[string]string, 10)
	fmt.Printf("清空后 m=%v\n", m)
}

// 3.遍历map
func Traversing(m map[string]string) {
	for k, v := range m {
		fmt.Printf("key=%v, value=%v\n", k, v)
	}
}

// 4.map排序
func MapSort() {
	m := map[int]int{
		10: 100, 1: 13, 4: 56, 8: 90,
	}
	
	// 将全部key取出存入切片
	var keys []int
	for k, _ := range m {
		keys = append(keys, k)
	}
	
	// 排序切片
	sort.Ints(keys)

	// 按排序后的顺序依次取出value
	for _, k := range keys {
		fmt.Printf("key=%v, value=%v\n", k, m[k])
	}
}

// 5.词频统计示例
func WordCount(str string) map[string]int {
	slice := strings.Fields(str)
	result := make(map[string]int)
	
	for _, val := range slice {
		if num, ok := result[val]; !ok {
			result[val] = 1
		} else {
			num++
			result[val] = num
		}
	}
	
	return result
}

func main() {
	// map是引用类型，能动态增长key-value对
	m := Create()
	fmt.Println()

	CRUD(m)
	fmt.Println()

	Traversing(m)
	fmt.Println()

	MapSort()
	fmt.Println()

	result := WordCount("I love my work and I love my family too")
	fmt.Println(result)
}