package main

import (
	"fmt"
	"sync/atomic"
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
	baseid := uint64(1000)
	tmp := atomic.AddUint64(&baseid, 1)
	fmt.Println(baseid, tmp)


}