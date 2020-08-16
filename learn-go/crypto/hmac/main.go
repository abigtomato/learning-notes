package main

import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
)

// 生成消息认证码
func GenerateHMac(plainText, key []byte) []byte {
    // 1.创建一个采用sha256作为底层hash接口、key作为密钥的HMAC算法的hash接口
	myHash := hmac.New(sha256.New, key)
    // 2.向hash中添加明文数据
	myHash.Write(plainText)
    // 3.计算hash结果
	return myHash.Sum(nil)
}

// 验证消息认证码
func VerifyHMac(plainText, key, hashText []byte) bool {
    // 1.创建一个采用sha256作为底层hash接口、key作为密钥的HMAC算法的hash接口
	myHash := hmac.New(sha256.New, key)
    // 2.向hash中添加明文数据
	myHash.Write(plainText)
	// 3.计算hmac
    hamcl := myHash.Sum(nil)
    // 4.比较两个hmac是否相同
	return hmac.Equal(hashText, hamcl)
}

func main() {
	src := []byte("Hello, Golang")
	key := []byte("Hello World")
	hashText := GenerateHMac(src, key)
	final := VerifyHMac(src, key, hashText)
	if final {
		fmt.Println("消息认证码认证成功!!!")
	} else {
		fmt.Println("消息认证码认证失败...")
	}
}
