package main

import (
	"fmt"
	"time"
	"strconv"
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

func main2() {
	data := make([]byte, 0, 512)
	a := []byte{1, 2, 3, 4}

	copy(data, a[:4])

	fmt.Println(data[0], len(data), cap(data))

	data = data[0:4]

	fmt.Println(data, len(data), cap(data))

}

func main() {
	count := 100000000
	idx := 0
	ts := time.Now()
	var b bool
	abc := []byte{49, 50, 51}

	for ; idx < count ;{

		b = string(abc) == "123"
		idx += 1
	}

	fmt.Println(time.Since(ts), b)

	idx = 0
	for ; idx < count ;{
		def, _ := strconv.ParseUint(string(abc), 10, 64)
		b = def == 123
		idx += 1
	}

	fmt.Println(time.Since(ts), b)


	i, _ := strconv.ParseUint(string(abc), 10, 64)
	j := string(abc)
	fmt.Println(i, j)

}
