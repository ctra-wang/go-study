package main

import (
	"fmt"
	"sync"
)

func main() {

	//slice := []int{}
	chan1 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {

		go func() {
			defer wg.Done()
			temp := i + 100
			chan1 <- temp
		}()
	}

	wg.Add(5)
	go func() {
		defer wg.Done()
		val := <-chan1
		fmt.Println(val)
	}()
	wg.Wait()
}
