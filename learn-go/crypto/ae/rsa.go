package main

import (
	"os"
	"fmt"
	"crypto/rand"
	"crypto/x509"
	"crypto/rsa"
	"encoding/pem"
)

// 生成RSA密钥对，并持久化到磁盘中
func GenerateRsaKey(keySize int) {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		fmt.Printf("rsa.GenerateKey Error: %v\n", err)
		return
	}

	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	block := pem.Block{
		Type: "rsa private key",
		Bytes: derText,
	}

	file, err := os.Create("./pem/private.pem")
	if err != nil {
		fmt.Printf("os.Create Error: %v\n", err)
		return
	}
	defer file.Close()
	pem.Encode(file, &block)

	publicKey := privateKey.PublicKey
	derStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		fmt.Printf("x509.MarshalPKIXPublicKey Error: %v\n", err)
		return
	}

	file, err = os.Create("./pem/public.pem")
	if err != nil {
		fmt.Printf("os.Create Error: %v\n", err)
		return
	}
	defer file.Close()
	pem.Encode(file, &pem.Block{
		Type: "rsa public key",
		Bytes: derStream,
	})
}

// 使用RSA公钥进行加密
func RSAEncrypt(plainText []byte, publicKeyFileName string) []byte {
	file, err := os.Open(publicKeyFileName)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		return nil
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("file.Stat Error: %v\n", err)
		return nil 
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)

	block, _ := pem.Decode(buf)
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKIXPublicKey Error: %v\n", err)
		return nil 
	}
	publicKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		fmt.Println("pubInterface.(*rsa.PublicKey) Error")
		return nil
	}

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		fmt.Printf("rsa.EncryptPKCS1v15 Error: %v\n", err)
		return nil
	}
	return cipherText
}

// 使用RSA私钥进行解密
func RSADecrypt(cipherText []byte, privateKeyFileName string) []byte {
	file, err := os.Open(privateKeyFileName)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		return nil
	}
	defer file.Close()
	
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("file.Stat Error: %v\n", err)
		return nil
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)

	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKCS1PrivateKey Error: %v\n", err)
		return nil
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		fmt.Printf("rsa.DecryptPKCS1v15 Error: %v\n", err)
		return nil
	}
	return plainText
}

func main() {
	GenerateRsaKey(1024)

	src := []byte("Golang才是世界上最好的语言")
	
	cipherText := RSAEncrypt(src, "./pem/public.pem")
	fmt.Printf("加密后: %s\n", cipherText)
	
	plainText := RSADecrypt(cipherText, "./pem/private.pem")
	fmt.Printf("解密后: %s\n", plainText)
}
