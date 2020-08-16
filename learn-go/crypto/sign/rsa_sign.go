package main

import (
	"fmt"
	"os"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"crypto/rand"
	"crypto/sha512"
	"encoding/pem"
)

// RSA签名
func SignatureRSA(plainText []byte, fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		return nil
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Printf("file.Stat Error: %v\n", err)
		return nil
	}
	buf := make([]byte, info.Size())
	file.Read(buf)

	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKCS1PrivateKey Error: %v\n", err)
		return nil
	}

	myHash := sha512.New()
	myHash.Write(plainText)
	hashText := myHash.Sum(nil)

	signText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashText)
	if err != nil {
		fmt.Printf("rsa.SignPKCS1v15 Error: %v\n", err)
		return nil
	}
	return signText
}

// RSA签名认证
func VerifyRSA(plainText, signText []byte, fileName string) bool {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		return false
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Printf("file.Stat Error: %v\n", err)
		return false
	}
	buf := make([]byte, info.Size())
	file.Read(buf)

	block, _ := pem.Decode(buf)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParsePKIXPublicKey Error: %v\n", err)
		return false
	}
	publicKey := pub.(*rsa.PublicKey)
	hashText := sha512.Sum512(plainText)
	
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashText[:], signText)
	if err != nil {
		fmt.Printf("rsa.VerifyPKCS1v15 Error: %v\n", err)
		return false
	}
	return true
}

func main() {
	src := []byte("Hello, Kubernetes")
	signText := SignatureRSA(src, "./rsa/private.pem")	
	bl := VerifyRSA(src, signText, "./rsa/public.pem")
	fmt.Println(bl)
}