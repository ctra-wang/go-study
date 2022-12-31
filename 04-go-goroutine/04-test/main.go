package main

import (
	"fmt"
	"time"
)

func MackeChan() chan int {
	chan1 := make(chan int)

	for i := 0; i < 5; i++ {

		go func(i int) {
			temp := i + 100
			chan1 <- temp
		}(i)
	}

	return chan1
}

func main() {

	chan1 := MackeChan()
	close(chan1)
	time.Sleep(time.Second * 1)
	for {
		res, ok := <-chan1
		if !ok {
			fmt.Println("关闭通道")
			break
		}
		fmt.Println(res)
	}

	//wg.Add(5)
	//go func() {
	//	defer wg.Done()
	//	val := <-chan1
	//	fmt.Println(val)
	//}()
	//wg.Wait()
}
