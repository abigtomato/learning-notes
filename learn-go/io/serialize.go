package main

import (
	"fmt"
	"encoding/json"
)

/* 
	序列化
 */

type Monster struct {
	Name 	 string  `json:"monster_name"`	// struct tag为序列化为json时指定key的名称
	Age 	 int 	 `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Sal 	 float64 `json:"monster_sal"`
	Skill 	 string  `json:"monster_skill"`
}

// 序列化结构体
func serialStruct() {
	monster := &Monster{
		Name: "albert",
		Age: 20,
		Birthday: "2018-10-19",
		Sal: 8000,
		Skill: "king",
	}
	var monster2 Monster

	if data, err := json.Marshal(monster); err != nil {
		fmt.Printf("err=%v\n", err)
	} else {
		fmt.Printf("结构体序列化=%v\n", string(data))
		if err := json.Unmarshal(data, &monster2); err != nil {
			fmt.Printf("err=%v\n", err)
		} else {
			fmt.Printf("结构体反序列化=%v\n", monster2)
		}
	}
}

// 序列化映射
func serialMap() {
	m := map[string]interface{}{
		"name": "albert",
		"age": 21,
		"score": 150,
	}
	var m2 map[string]interface{}

	if data, err := json.Marshal(m); err != nil {
		fmt.Printf("err=%v\n", err)
	} else {
		fmt.Printf("映射序列化=%v\n", string(data))
		// 反序列化map不需要make分配内存，因为make操作已经被封装到Unmarshal函数中了
		if err := json.Unmarshal(data, &m2); err != nil {
			fmt.Printf("err=%v\n", err)
		} else {
			fmt.Printf("映射反序列化=%v\n", m2)
		}
	}
}

// 序列化切片
func serialSlice() {
	var slice = make([]map[string]interface{}, 0, 10)
	for i := 0; i < 10; i++ {
		slice = append(slice, map[string]interface{}{
			"name": "albert",
			"age": 20 + i,
			"score": 150 + i * 100,
		})
	}
	var slice2 []map[string]interface{}

	if data, err := json.Marshal(slice); err != nil {
		fmt.Printf("err=%v\n", err)
	} else {
		fmt.Printf("切片序列化=%v\n", string(data))
		if err := json.Unmarshal(data, &slice2); err != nil {
			fmt.Printf("err=%v\n", err)
		} else {
			fmt.Printf("切片反序列化=%v\n", slice2)
		}
	}
}

func serEncoding() {
	types := []string{"Hadoop", "Spark", "Docker"}

	var buffer bytes.Buffer
	enc := gob.NewEncode(&buffer)
	enc.Encode(typs)
	data := buffer.Bytes()
	fmt.Printf("Encodeing Data: %v\n", data)

	dec := gob.NewDecode(bytes.NewReader(data))
	var dtypes []string
	dec.Decode(dtypes)
	fmt.Printf("Decoding Data: %v\n", dtypes)
}

func main() {
	serialStruct()
	fmt.Println()

	serialMap()
	fmt.Println()

	serialSlice()
	fmt.Println()
}