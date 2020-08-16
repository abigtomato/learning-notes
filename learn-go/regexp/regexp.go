package main

import (
	"fmt"
	"regexp"
)

const text = `my email is ccmouse@gmail.com@abc.com`

const emails = `
my email is ccmouse@gmail.com@abc.com
email1 is abc@def.org
email2 is kkk@163.com
email2 is kkk@abc.com.cn`

const url = `http://album.zhenai.com/u/10087567`

const html = `
<head>
	<title>首页 - Go语言中文网 - Golang中文社区</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
	<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
	<meta charset="utf-8">
	
	<link rel="shortcut icon" href="https://static.studygolang.com/img/favicon.ico">
	<link rel="apple-touch-icon" type="image/png" href="https://static.studygolang.com/static/img/logo2.png">
	
	<meta name="keywords" content="Go语言中文网,Go,Golang,Go语言,主题,资源,文章,图书,开源项目">
	<meta name="description" content="Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
	<meta name="author" content="polaris <polaris@studygolang.com>">
	
	<link rel="canonical" href="https://studygolang.com/" />
	<link rel="stylesheet" href="https://cdn.staticfile.org/bootswatch/3.2.0/css/cosmo/bootstrap.min.css">
	<link rel="stylesheet" href="https://cdn.staticfile.org/font-awesome/4.7.0/css/font-awesome.min.css">
	<link rel="stylesheet" href="https://static.studygolang.com/static/dist/css/sg_libs.min.css?v=20180305"/>
	<link rel="stylesheet" href="https://static.studygolang.com/static/dist/css/sg_styles.min.css?v=2018110320"/>

	<script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
	<script src="https://oss.maxcdn.com/libs/respond.js/1.3.0/respond.min.js"></script>
	<script>
		def fib(total):
			if total == 1 or total == 2:
				return 1
			else:
				return fib(total - 1) + fib(total - 2)
	</script>
</head>
`

func htmlParse() {
	// (?s) 是正则的模式修饰符，即Singleline单行模式，表示更改“.”的含义，使其匹配范围包含\n
	// (.*?) 是单元分组，“.”匹配任意字符，“*?”表示前一个字符重复>=0次匹配(非贪婪)
	re := regexp.MustCompile(`<script>(?s:(.*?))</script>`)
	matchs := re.FindAllStringSubmatch(html, -1)
	for _, match := range matchs {
		fmt.Printf("%v\n", match)
	}
}

func main() {
	// regexp.MustCompile: 将正则表达式解析成Go编译器识别的结构体格式
	// 参数: 正则表达式字符串，使用反引号(原生字符串，不对转义字符做任何转换)
	// 返回值: 解析后的结构体指针 *Regexp
	re := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	
	match := re.FindString(text)
	fmt.Printf("[类型: %T, 内容: %v]\n", match, match)

	matchSlice := re.FindAllString(emails, -1)
	fmt.Printf("[类型: %T, 内容: %v]\n", matchSlice, matchSlice)

	re = regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	// re.FindAllStringSubmatch: 根据正则解析目标字符串返回匹配结果
	// 参数1: 待解析的字符串
	// 参数2: 匹配的次数，-1代表匹配所有
	// 返回值: 返回匹配成功的 [][]string(二维数组中的每一个元素即一维数组，就是一个匹配的结果)
	matchArr := re.FindAllStringSubmatch(emails, -1)
	for _, m := range matchArr {
		fmt.Printf("[类型: %T, 内容: %v]\n", m, m)
	}

	re = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
	matchSlice = re.FindStringSubmatch(url)
	fmt.Printf("[类型: %T, 内容: %v]\n", matchSlice, matchSlice)

	htmlParse()
}