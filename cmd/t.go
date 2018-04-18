package main

import (
	"fmt"
	"time"
	"strconv"
	"rpcpool"
	"net/rpc"
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

func main3 () {
	count := 1
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



	tt := make(map[string]int)
	fmt.Println(tt, len(tt))

	var aa []byte

	aa = nil

	if v, ok := tt[string(aa)]; ok {
		fmt.Println(v, ok)
	}


}

var addr string = "192.168.1.2:1235"

var pool rpcpool.Pool

func main4 () {
	ts := time.Now()
	rpc.DialHTTP("tcp", addr)
	fmt.Println(time.Since(ts))


	pool = rpcpool.Pool{
		MaxIdle:     100,
		MaxActive:   10000,
		Dial: func() (rpcpool.Conn, error) {
			return rpc.DialHTTP("tcp", addr)
		},
	}

	allcount := 1000

	for {
		ch := make(chan time.Duration, allcount)
		all_time := float64(0)
		for idx := 0; idx < 1000; idx += 1 {
			go t(idx,  ch)
		}
		max := float64(-1)
		for idx := 0; idx < allcount; idx += 1 {
			x := <-ch
			tmp := x.Seconds()
			if max < tmp {
				max = tmp
			}
			all_time += x.Seconds()
		}

		fmt.Println(all_time, all_time/float64(allcount), max)
		fmt.Println("--------------------------------", pool.ActiveCount(), pool.IdleCount())
		time.Sleep(5 * time.Second)

	}
	fmt.Println("DONE.")

	for {
		time.Sleep(3 * time.Second)
	}

	defer pool.Close()

}


func t(idx int, ch chan time.Duration) {
	ts := time.Now()
	conn := pool.Get()
	ch <- time.Since(ts)
	defer conn.Close()
	time.Sleep(time.Millisecond * 1)

}

func t1(idx int, ch chan time.Duration) {
	ts := time.Now()
	conn := pool.Get()
	ch <- time.Since(ts)
	defer conn.Close()
	time.Sleep(time.Millisecond * 1)

}

func main() {
	allcount := 1000
	pSize := 10
	clis := make([]*rpc.Client, pSize)
	status := make(p[])

	for idx:=0;idx<pSize;idx+=1 {
		cli, _ := rpc.DialHTTP("tcp", addr)
		clis[idx] = cli
	}

	fmt.Println("init pool done.")

	ch := make(chan time.Duration, allcount)
	all_time := float64(0)
	for idx := 0; idx < 1000; idx += 1 {
		go func(idx int,  ch chan time.Duration) {

		}(idx, ch)
	}
	max := float64(-1)
	for idx := 0; idx < allcount; idx += 1 {
		x := <-ch
		tmp := x.Seconds()
		if max < tmp {
			max = tmp
		}
		all_time += x.Seconds()
	}

	fmt.Println(all_time, all_time/float64(allcount), max)
}