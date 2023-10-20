package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	index int32
	block = make(chan struct{}, 10)
)

func readDB() string {
	atomic.AddInt32(&index, 1)
	fmt.Printf("并发调度---index %d \n", atomic.LoadInt32(&index))
	time.Sleep(100 * time.Millisecond)
	atomic.AddInt32(&index, -1)
	return "OK"
}

func handler() {
	block <- struct{}{}
	readDB()
	<-block
	return
}

func main() {
	for i := 0; i < 100; i++ {
		go handler()
	}
	time.Sleep(3 * time.Second)
}
