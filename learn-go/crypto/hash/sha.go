package main

import (
	"fmt"
	"os"
	"io"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func GetSha1(fileName string) (result string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		return
	}
	defer file.Close()

	myHash := sha1.New()
	num, err := io.Copy(myHash, file)
	if err != nil {
		fmt.Printf("io.Copy Error: %v\n", err)
		return
	}
	fmt.Println("文件大小: ", num)

	result = hex.EncodeToString(myHash.Sum(nil))
	fmt.Println("sha1: ", result)
	return
}

func GetSha256(fileName string) (result string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		return
	}
	defer file.Close()

	myHash := sha256.New()
	num, err := io.Copy(myHash, file)
	if err != nil {
		fmt.Printf("io.Copy Error: %v\n", err)
		return
	}
	fmt.Println("文件大小: ", num)

	result = hex.EncodeToString(myHash.Sum(nil))
	fmt.Println("sha1: ", result)
	return
}

func GetSha512(fileName string) (result string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		return
	}
	defer file.Close()

	myHash := sha512.New()
	num, err := io.Copy(myHash, file)
	if err != nil {
		fmt.Printf("io.Copy Error: %v\n", err)
		return
	}
	fmt.Println("文件大小: ", num)

	result = hex.EncodeToString(myHash.Sum(nil))
	fmt.Println("sha1: ", result)
	return
}

func main() {
	res := sha256.Sum256([]byte("Hello, Golang"))
	fmt.Println(res)

	hash := sha256.New()
	hash.Write([]byte("Hello, Python"))
	hash.Write([]byte("Hello, C++"))
	hash.Write([]byte("Hello, Scala"))
	ret := hash.Sum(nil)	// 通过sha256哈希后的散列是32字节
	str := hex.EncodeToString(ret)	// 通过转码为16进制数后变为64个字节
	fmt.Println(len(str), str)
}
