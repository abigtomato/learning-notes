package main

import "fmt"

/* 哈希表 */

 type EmpNode struct {
	Id 		int
	Name 	string
	Next 	*EmpNode
}

type EmpLink struct {
	Head *EmpNode
}

func (this *EmpLink) Insert(emp *EmpNode) {
	if this.Head == nil {
		this.Head = emp
		return 
	}
	
	var cur, pre *EmpNode
	cur = this.Head
	pre = nil

	for cur != nil {
		if cur.Id >= emp.Id {
			break
		}
		pre = cur
		cur = cur.Next
	}
	pre.Next = emp
	emp.Next = cur
}

func (this *EmpLink) ShowList() {
	if this.Head == nil {
		fmt.Println("linkedlist empty")
		return
	}
	
	cur := this.Head
	for cur != nil {
		fmt.Printf("[Id: %v, Name: %v, Next: %p]--->", cur.Id, cur.Name, cur.Next)
		cur = cur.Next
	}
}

func (this *EmpLink) FindEmpById(id int) *EmpNode {
	if this.Head == nil {
		fmt.Println("linkedlist empty")
		return nil
	}

	cur := this.Head
	for cur != nil {
		if cur.Id == id {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) Insert(emp *EmpNode) {
	insertNo := this.HashFun(emp.Id)
	this.LinkArr[insertNo].Insert(emp)
}

func (this *HashTable) HashFun(id int) int {
	return id % (len(this.LinkArr) - 1)
}

func (this *HashTable) ShowList() {
	for index, value := range this.LinkArr {
		fmt.Printf("linkedlist %v number: ", index)
		value.ShowList()
		fmt.Println()
	}
}

func (this *HashTable) FindEmpById(id int) *EmpNode {
	findNo := this.HashFun(id)
	return this.LinkArr[findNo].FindEmpById(id)
}

func main() {
	hashTable := &HashTable{}
	
	for i := 6; i <= 54; i += 6 {
		emp := &EmpNode{
			Id: i,
			Name: fmt.Sprintf("albert %v", i),
		}
		hashTable.Insert(emp)
	}

	hashTable.ShowList()

	res := hashTable.FindEmpById(42)
	fmt.Printf("result: %v\n", res.Name)
}