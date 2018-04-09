package main

import (
	"fmt"
)

func main() {
	a := []byte{}
	fmt.Println(a == nil)
	fmt.Println(len(a))

	for idx, item := range a {
		fmt.Println(idx, item)
	}

	fmt.Println("------------")
	var b []byte
	fmt.Println(b == nil)
	fmt.Println(len(b))

	for idx, item := range b {
		fmt.Println(idx, item)
	}
	fmt.Println("------------")
}
