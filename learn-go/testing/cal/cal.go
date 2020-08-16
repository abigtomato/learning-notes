package cal

// 需要测试的函数
func AddUpper(n int) int {
	num := 0
	for i := 1; i <= 10; i++ {
		num += i
	}
	return num
}

// 需要测试的函数
func GetSub(n1, n2 int) int {
	return n1 - n2
}

// 需要测试的函数
func Fib(total int) int {
	if total == 1 || total == 2 {
		return 1
	} else {
		return Fib(total - 1) + Fib(total - 2)
	}
}