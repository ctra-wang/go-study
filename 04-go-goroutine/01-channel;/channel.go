package main

import (
	"sync"
)

func main() {
	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		c <- `foo`
	}()

	go func() {
		defer wg.Done()

		//time.Sleep(time.Second * 1)
		println(`Message: ` + <-c)
	}()

	wg.Wait()
}
