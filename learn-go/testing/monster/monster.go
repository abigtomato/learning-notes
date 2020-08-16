package monster

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func (this *Monster) Store(path string) bool {
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return false
	}

	err = ioutil.WriteFile(path, data, 0666)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return false
	}

	return true
}

func (this *Monster) ReStore(path string) bool {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return false
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return false
	}

	return true
}