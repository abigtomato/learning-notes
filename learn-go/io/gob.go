package main

import (
	"fmt"
	"log"
	"bytes"
	"encoding/gob"
)

/*  */

// 测试结构
type Person struct {
	Name string
	Age int
}

// gob序列化
func Serialize(data *Person) []byte {
	/* func NewEncoder(w io.Writer) *Encoder
	   NewEncoder返回一个将编码后数据写入w的*Encoder */
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)	// 获取编码器
	
	/* func (enc *Encoder) Encode(e interface{}) error
		Encode方法将e编码后发送，并且会保证所有的类型信息都先发送 */
	err := encoder.Encode(data)	// 编码
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}

// gob反序列化
func DeSerialize(src []byte, dest *Person) {
	/* func NewDecoder(r io.Reader) *Decoder
	   函数返回一个从r读取数据的*Decoder，如果r不满足io.ByteReader接口，则会包装r为bufio.Reader */
	decoder := gob.NewDecoder(bytes.NewReader(src))

	/* func (dec *Decoder) Decode(e interface{}) error
	   Decode从输入流读取下一个之并将该值存入e。如果e是nil，将丢弃该值；否则e必须是可接收该值的类型的指针。
	   如果输入结束，方法会返回io.EOF并且不修改e（指向的值） */
	err := decoder.Decode(dest)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	person := Person{"Albert", 22}
	encode := Serialize(&person)
	fmt.Printf("encode: %v\n", encode)
	
	var decode Person
	DeSerialize(encode, &decode)
	fmt.Printf("decode: %v\n", decode)
}