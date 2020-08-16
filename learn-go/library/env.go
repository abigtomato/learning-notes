package main

import (
	"os"
	"fmt"
)

func main() {
	err := os.Setenv("TEST", "scar_test")
	if err != nil {
		fmt.Println(err)
	}
}
