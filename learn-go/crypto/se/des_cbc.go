package main

import (
	"fmt"
	"bytes"
	"crypto/des"
	"crypto/cipher"
)

// 填充函数
func paddingLastGroup(plainText []byte, blockSize int) []byte {
	padNum := blockSize - len(plainText) % blockSize
	char := []byte{byte(padNum)}
	newPlain := bytes.Repeat(char, padNum)
	return append(plainText, newPlain...)
}

// 去除填充数据
func unPaddingLastGrooup(plainText []byte) []byte {
	length := len(plainText)
	lastChar := plainText[length - 1]
	number := int(lastChar)
	return plainText[: length - number]
}

// 使用DES算法，CBC分组模式加密
func desEncrypt(plainText, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Printf("des.NewCipher Error: %v", err)
		return nil
	}

	newText := paddingLastGroup(plainText, block.BlockSize())
	iv := []byte("12345678")
	blockModel := cipher.NewCBCEncrypter(block, iv)

	cipherText := make([]byte, len(newText))
	blockModel.CryptBlocks(cipherText, newText)

	return cipherText
}

// 使用DES算法解密
func desDecrypt(cipherText, key []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Printf("des.NewCipher Error: %v\n", err)
		return nil
	}
	
	iv := []byte("12345678")
	blockModel := cipher.NewCBCDecrypter(block, iv)
	blockModel.CryptBlocks(cipherText, cipherText)
	
	return unPaddingLastGrooup(cipherText)
}

func main() {
	key := []byte("1234abcd")
	src := []byte("PHP不是世界上最好的语言")
	cipherText := desEncrypt(src, key)
	fmt.Printf("加密后: %s\n", cipherText)
	plainText := desDecrypt(cipherText, key)
	fmt.Printf("解密后: %s\n", plainText)
}
