package cal

import "testing"

// Testing框架调用过程:
// 1.cal_test.go导入cal.go，TestAddUpper()函数调用需要测试的函数AddUpper()
// 2.testing将cal_test.go文件导入
// 3.testing调用TestAddUpper()函数
func TestAddUpper(t *testing.T) {
	const RES = 55
	res := AddUpper(10)

	if res != RES {
		t.Fatalf("AddUpper(10) 执行错误, 期望值=%v, 实际值=%v\n", RES, res)
	} else {
		t.Logf("AddUpper(10) 执行正确 ...")
	}
}

// Testing测试框架使用注意事项:
// 1.测试用例文件名必须以_test.go结尾
// 2.测试用例函数必须以Test开头，一般来说就是Test+被测试的函数名
// 3.测试用例函数的形参必须固定是*testing.T类型
// 4.一个测试文件中可以包含多个测试用例函数 
// 5.运行测试用例的命令: 
// 		cmd> go test	如果运行正确无日志输出，错误会输出日志
// 		cmd> go test -v	无论运行正确或是错误，都会输出日志
// 6.出现错误使用t.Fatalf()格式化输出错误信息，并退出程序
// 7.使用t.Logf()输出相应的日志
// 8.单独运行一个测试文件 go test -v cal_test.go cal.go
// 9.单独运行一个测试用例 go test -v -test.run TestAddUpper
func TestGetSub(t *testing.T) {
	const RES = 10
	res := GetSub(18, 8)

	if res != RES {
		t.Fatalf("GetSub(18, 8): 执行错误, 期望值: %v, 实际值: %v\n", RES, res)
	} else {
		t.Logf("GetSub(18, 8) 执行正确 ...")
	}
}

// 表格驱动测试
func TestFib(t *testing.T) {
	tests := []struct{
		total int
		result int
	}{
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
	}

	for _, test := range tests {
		ret := Fib(test.total)
		if ret != test.result {
			t.Errorf("Fib(%d): 执行错误, 期望值: %d, 实际值: %d", test.total, test.result, ret)
		} else {
			t.Logf("Fib(%d) 执行正确 ...", test.total)
		}
	}
}

// Testing框架的性能测试:
// 1.运行测试用例的命令:
// 		cmd> go test -bench .
// 2.参*testing.B提供性能测试的方法
func BenchmarkSubstr(b *testing.B) {
	total := 20
	result := 6765

	// 循环次数由Testing提供的b.N控制
	for i := 0; i < b.N; i++ {
		if ret := Fib(total); ret != result {
			b.Errorf("Fib(%d): 执行错误, 期望值: %d, 实际值: %d", total, result, ret)
		}
	}
}