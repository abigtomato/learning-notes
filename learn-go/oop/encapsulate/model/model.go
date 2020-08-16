package model

// 封装1: 结构体名首字母小写，表示包内私有，只允许包外使用工厂函数创建实例
type person struct {
	// 封装2: 属性首字母小写，表示为私有，包外不可访问，只允许包外通过对外开放的方法操作
	name string
	age int
	sal float64
}

// 封装3: 提供工厂函数，只能通过工厂函数创建结构体的实例
func NewPerson(name string, age int, sal float64) *person {
	return &person {name, age, sal}
}

// 封装4: 提供包外可见的对结构体属性进行操作的方法
func (p *person) SetAge(age int) {
	if age >= 150 || age <= 0 {
		return
	}
	p.age = age  
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal > 30000 || sal < 4000 {
		return
	} 
	p.sal = sal
}

func (p *person) GetSal() float64 {
	return p.sal
}