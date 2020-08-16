package main

import "fmt"

/*
	1.位运算

	原码（二进制码），反码，补码:
		1. 二进制的最高位是符号位，0表示正数，1表示负数；
		2. 正数的原码，反码，补码都一样：
			1 => 原码[0000 0001], 反码[0000 0001], 补码[0000 0001]；
		3. 负数的反码等于原码符号位不变其他位取反，补码等于其反码加1：
			-1 => 原码[1000 0001], 反码[1111 1110], 补码[1111 1111]；
		4. 0的反码，补码都是0；
		5. 计算机运算时，都是以补码的方式运算的。
 */
func BitOperation() {
	// 2的补码 => [0000 0010]
	// 3补码 => [0000 0011]
	// 2&3（全1为1，有0则0）=> [0000 0010] => 2 
	fmt.Printf("按位与[&]: 二进制 = %b, 十进制 = %d\n", 2&3, 2&3)

	// 2补码 => [0000 0010]
	// 3补码 => [0000 0011]
	// 2|3（有1则1，全0为0）=> [0000 0011] => 3
	fmt.Printf("按位或[|]: 二进制 = %b, 十进制 = %d\n", 2|3, 2|3)

	// 2补码 => [0000 0010]
	// 3补码 => [0000 0011]
	// 2^3（0和1为1，全0全1为0）=> [0000 0001] => 1 
	fmt.Printf("按位异或[^]: 二进制 = %b, 十进制 = %d\n", 2^3, 2^3)

	// -2原码[1000 0010]
	// -2反码[1111 1101]
	// -2补码[1111 1110]
	//  2补码[0000 0010]
	// -2^2 => 1111 1100 (补码)
	// -2^2 => 1111 1011 (反码)
	// -2^2 => 1000 0100 => -4 (原码)
	fmt.Printf("负数按位异或[^]: 二进制 = %b, 十进制 = %d\n", -2^2, -2^2)

	// 1补码[0000 0001]
	// 1>>2 => [0000 0000]（低位溢出，符号位补溢出的高位）
	fmt.Printf("右移运算[>>]: %b\n", 1>>2)

	// 1补码[0000 0001]
	// 1<<2 => 0000 0100（符号位不变，低位补0）
	fmt.Printf("左移运算[<<]: %b\n", 1<<2)
}

// 2.二进制转十进制
func Binary2Decimal() {
	// 从最低位开始，将每位数据提取出来，乘以 2^位数-1次方 后求和
	// 1011 = (1 * 2^0) + (1 * 2^1) + (0 * 2^2) + (1 * 2^3) = 1 + 2 + 0 + 8 = 11
	var bnum int = 11
	
	fmt.Printf("二进制: %b => 十进制: %d\n", bnum, bnum)
}

// 3.十进制转二进制
func Decimal2Binary() {
	// 将数循环除2取余，直到商为0，之后将每步得到的余数倒置连接
	// 56 / 2 = 28, 56 % 2 = 0
	// 28 / 2 = 14, 28 % 2 = 0
	// 14 / 2 = 7, 14 % 2 = 0
	// 7 / 2 = 3, 7 % 2 = 1
	// 3 / 2 = 0, 3 % 2 = 1
	// 余数倒置：11000
	var bnum int = 56
	
	fmt.Printf("十进制: %d => 二进制: %b\n", bnum, bnum)
}

// 4.八进制转十进制
func Octal2Decimal() {
	// 从最低位开始，将每位数据提取出来，乘以 8^位数-1次方 后求和
	// 0123 = 3 * 8^0 + 2 * 8^1 + 1 * 8^2 + 0 * 8^3 = 3 + 16 + 64 = 83
	var onum int = 0123
	
	fmt.Printf("八进制: %o => 十进制: %d\n", onum, onum)
}

// 5.十进制转八进制
func Decimal2Octal() {
	// 将该数不断除8取余，直到商为0，之后将每步得到的余数倒置连接
	// 156 / 8 = 19, 156 % 8 = 4
	// 19 / 8 = 2, 19 % 8 = 3
	// 2 / 8 = 0, 2 % 8 = 2  
	// 余数倒置：234
	var onum int = 156
	
	fmt.Printf("十进制: %d => 八进制: %o\n", onum, onum)
}

// 6.十六进制转十进制
func Hexadecimal2Decimal() {
	// 从最低位开始，将每位数据提取出来，乘以 16^位数-1次方 后求和
	// 0x34A = 10 * 16^0 + 4 * 16^1 + 3 * 16^2 = 10 + 64 + 768 = 842
	var xnum int = 0x34A
	
	fmt.Printf("十六进制: %x => 十进制: %d\n", xnum, xnum)
}

// 7.十进制转十六进制
func Decimal2Hexadecimal() {
	// 将该数不断除16取余，直到商为0，之后将每步得到的余数倒置连接
	// 356 / 16 = 22, 356 % 16 = 4
	// 22 / 16 = 1, 22 % 16 = 6
	// 1 / 16 = 0, 1 % 16 = 1
	// 余数倒置：164
	var xnum int = 356
	
	fmt.Printf("十进制进制: %d => 十六进制: %x\n", xnum, xnum)
}

// 8.二进制转八进制
func Binary2Octal() {
	// 将二进制数每三位一组（从低位开始组合，不足三位的划分为单独一组）
	// 每组各自转换为10进制数，各组的结果组合起来就是最终结果
	// 11010101 => 11, 010, 101 => 3, 2, 5 => 325
	num := 0325
	
	fmt.Printf("二进制: %b => 八进制: %o\n", num, num)
}

// 9.二进制转十六进制
func Binary2Hexadecimal() {
	// 将二进制数每四位一组（从低位开始组合，不足四位的划分为单独一组）
	// 每组转换为10进制数，每组的结果组合起来就是最终结果
	// 11010101 => 1101, 0101 => D5
	num := 0xD5
	
	fmt.Printf("二进制: %b => 十六进制: %x\n", num, num)
}

// 10.八进制转二进制
func Octal2Binary() {
	// 将八进制数的每一位，按十进制转二进制的规则（除二取余）转成对应的3位二进制数
	// 将每一位转换的结果组合起来就是最终结果
	// 237 => 10, 011, 111 => 10011111
	num := 0237
	
	fmt.Printf("八进制: %o => 二进制: %b\n", num, num)
}

// 11.十六进制转二进制
func Hexadecimal2Binary() {
	// 将十六进制数的每一位，按十进制转二进制的规则（除二取余）转成对应的4位二进制数
	// 将每一位转换的结果组合起来就是最终结果
	// 237 => 10, 0011, 0111 => 1000110111
	num := 0x237
	
	fmt.Printf("十六进制: %x => 二进制: %b\n", num, num)
}

func main() {
	BitOperation()
	fmt.Println()

	Binary2Decimal()
	Decimal2Binary()
	fmt.Println()

	Octal2Decimal()
	Decimal2Octal()
	fmt.Println()

	Hexadecimal2Decimal()
	Decimal2Hexadecimal()
	fmt.Println()

	Binary2Octal()
	Binary2Hexadecimal()
	fmt.Println()

	Octal2Binary()
	Hexadecimal2Binary()
}