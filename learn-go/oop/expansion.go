package main

import "fmt"

type Monkey struct {	
	Name string
	Age int
}

// 继承主要是解决代码的复用性和可维护性
type LittleMonkey struct {
	Monkey
}

func (this *Monkey) climbing() {
	fmt.Printf("%v生来就会climbing\n", this.Name)
}

// 接口主要是设计规范，让其他自定义类型去遵守规范实现扩展功能
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
} 

// 为LittleMonkey扩展功能特性
func (this *LittleMonkey) Flying() {
	fmt.Printf("%v通过学习学会了Flying\n", this.Name)
}

func (this *LittleMonkey) Swimming() {
	fmt.Printf("%v通过学习学会了Swimming\n", this.Name)
}

func main() {
	monkey := &LittleMonkey{
		Monkey: Monkey{Name: "lily", Age: 3},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}