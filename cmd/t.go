package main

import (
	"fmt"
)

func main1() {
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

func main() {
	data := make([]byte, 0, 512)
	a := []byte{1, 2, 3, 4}

	copy(data, a[:4])

	fmt.Println(data[0], len(data), cap(data))

	data = data[0:4]

	fmt.Println(data, len(data), cap(data))

}
