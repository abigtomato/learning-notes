package main

import (
	"fmt"
	"os"
	"math/big"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/sha1"
	"encoding/pem"
)

func GenerateEccKey() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	if err != nil {
		fmt.Printf("ecdsa.GenerateKey Error: %v\n", err)
		return
	}

	derText, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		fmt.Printf("x509.MarshalECPrivateKey Error: %v\n", err)
		return
	}

	file, err := os.Create("./ecc/private.pem")
	if err != nil {
		fmt.Printf("os.Create Error: %v\n", err)
		return
	}
	defer file.Close()
	pem.Encode(file, &pem.Block{
		Type: "ecdsa private key",
		Bytes: derText,
	})

	publicKey := privateKey.PublicKey
	derText, err = x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		fmt.Printf("x509.MarshalPKIXPublicKey Error: %v\n", err)
		return
	}

	file, err = os.Create("./ecc/public.pem")
	if err != nil {
		fmt.Printf("os.Create Error: %v\n", err)
		return
	}
	defer file.Close()
	pem.Encode(file, &pem.Block{
		Type: "ecdsa public key",
		Bytes: derText,
	})
}

func EccSignature(plainText []byte, privFile string) (rText, sText []byte) {
	file, err := os.Open(privFile)
	if err != nil {
		fmt.Printf("os.Open Error: %v\n", err)
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		fmt.Printf("file.Stat Error: %v\n", err)
		return
	}
	buf := make([]byte, info.Size())
	file.Read(buf)

	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("x509.ParseECPrivateKey Error: %v\n", err)
		return
	}

	hashText := sha1.Sum(plainText)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashText[:])
	if err != nil {
		fmt.Printf("ecdsa.Sign Error: %v\n", err)
		return
	}

	rText, err = r.MarshalText()
	if err != nil {
		fmt.Printf("r.MarshalText Error: %v\n", err)
		return
	}
	sText, err = s.MarshalText()
	if err != nil {
		fmt.Printf("s.MarshalText Error: %v\n", err)
		return
	}
	return
}

func EccVerify(plainText, rText, sText []byte, pubFile string) bool {
	file, err := os.Open(pubFile)
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
	publicKey := pub.(*ecdsa.PublicKey)

	hashText := sha1.Sum(plainText)
	var r, s big.Int
	r.UnmarshalText(rText)
	s.UnmarshalText(sText)

	return ecdsa.Verify(publicKey, hashText, r, s)
}

func main() {
	GenerateEccKey()
}
