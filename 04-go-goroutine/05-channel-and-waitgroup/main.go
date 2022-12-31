package main

import (
	"fmt"
	"sync"
)

func main() {

	chan1 := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(1 / 2)
		if 1/2 == 1 {
			chan1 <- 0
		} else {
			chan1 <- 1
		}
	}()

	wg.Wait()
	close(chan1)

	res := <-chan1
	if res == 0 {
		fmt.Println("------------")
	} else {
		fmt.Println("++++++++++++")

	}
}
