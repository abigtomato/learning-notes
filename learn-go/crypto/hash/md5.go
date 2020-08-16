package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

func GetMD5_01(src []byte) string {
	// 1.直接通过sum函数计算数据的md5
    res := md5.Sum(src)
	fmt.Println(res)
	
    // 2.将md5结果格式化成16进制格式的字符串
	resStr := fmt.Sprintf("%x", res)
	fmt.Println(resStr)
	
    // 3.通过hex.EncodeToString函数将md5结果格式化成16进制格式的字符串
	resStr = hex.EncodeToString(res[:])
	fmt.Println(resStr)

	return resStr
}

func GetMD5_02(src ...[]byte) string {
	myHash := md5.New()
	for _, v := range src {
		myHash.Write(v)
	}
	
	result := myHash.Sum(nil)
	fmt.Println(result)

	resStr := fmt.Sprintf("%x", result)
	fmt.Println(resStr)

	resStr = hex.EncodeToString(result)
	fmt.Println(resStr)

	return resStr
}

func main() {
	GetMD5_01([]byte("Hello, SparkMLlib"))
	GetMD5_02([]byte("Hello, Docker"), []byte("Hello, kubernetes"), []byte("Hello, Kafka"))
}