package main

import (
	"fmt"
	"go_code/learn-go/encapsulate/model"
)

func main() {
	person := model.NewPerson("albert", 20, 4000)
	person.SetSal(5000)
	fmt.Println(person.GetSal())
}