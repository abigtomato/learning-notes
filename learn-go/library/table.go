package main

import (
	"fmt"
	"github.com/modood/table"
)

// 表格打印
type House struct {
	Name  	string
	Sigil 	string
	Motto 	string
}  

func main() {
	hs := []House{
		{"Stark", "direwolf", "Winter is coming"},
		{"Targaryen", "dragon", "Fire and Blood"},
		{"Lannister", "lion", "Hear Me Roar"},
	}

	table.Output(hs)

	s := table.Table(hs)
	fmt.Println(s)
}
