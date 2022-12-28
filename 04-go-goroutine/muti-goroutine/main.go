package main

import (
	"fmt"
	"time"
)

func main() {
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("xxxxxxxxxxxx")
			}
		}()
		//wg.Done()
		panic("------")
	}()

	//wg.Wait()
	time.Sleep(1 * time.Second)

	fmt.Println("over")

}
