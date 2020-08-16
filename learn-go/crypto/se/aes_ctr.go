package main

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
)

// 使用AES算法，CTR模式进行加解密
func aesCrypto(text, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("aes.NewCipher Error: %v\n", err)
		return nil
	}

	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(text, text)

	return text
}

func main() {
	key := []byte("12345678abcdefgh")
	src := []byte("Golang才是世界上最好的语言")
	cipherText := aesCrypto(src, key)
	fmt.Printf("加密后: %s\n", cipherText)
	plainText := aesCrypto(cipherText, key)
	fmt.Printf("解密后: %s\n", plainText)
}