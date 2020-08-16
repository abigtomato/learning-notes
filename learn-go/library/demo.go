package main

import (
	"fmt"
	"time"
	"math/rand"
)

func LoadReqId() string {
	s := "0-a48666442395560"
	date := time.Now().UnixNano() / 1e6
	keygen := Keygen("0123456789abcedf", 128)
	str := fmt.Sprintf("%v-%v-%v", s, date, keygen)
	return str[0:128]
}

func Keygen(str string, num int) string {
	rand.Seed(time.Now().Unix())
	chars := make([]byte, num)

	for index := range chars {
		chars[index] = str[rand.Intn(len(str))]
	}

	return string(chars)
}

func main() {
	res := LoadReqId()
	fmt.Println(res)
}
