package main

import (
	"os"
	"fmt"
	"strconv"
	"html/template"
)

/*
	模板引擎的使用
 */

type SearchResult struct {
	Hits int
	Start int
	Items []Item
}

type Item struct {
	Url string
	Payload User
}

type User struct {
	Id int
	Name string
	Age int
	Sex string
	Height int
	Weight int
}

func main() {
	// template.Must 检查模板语法
	// template.ParseFiles 指定模板文件
	template := template.Must(template.ParseFiles("./template.html"))

	out, err := os.Create("./template_example.html")
	page := SearchResult{
		Hits: 100,
		Start: 10,
		Items: make([]Item, 10),
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, Item{
			Url: fmt.Sprintf("http://www.imooc%d.com", i),
			Payload: User{
				Id: i,
				Name: "albert" + strconv.Itoa(i),
				Age: i * 10,
				Sex: "man",
				Height: 180,
				Weight: 140,
			},
		})
	}
	
	// template.Execute 开始通过模板渲染页面
	// 参数需要一个writer和一个页面填充数据
	err = template.Execute(out, page)
	if err != nil {
		panic(err)
	}
}