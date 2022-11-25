package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func thread(i int) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
}
