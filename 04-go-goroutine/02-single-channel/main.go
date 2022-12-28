package main

import "fmt"

// Producer 返回一个通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
func Producer() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				fmt.Println("当前有效数值i:", i)
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func Consumer(ch chan int) int {
	sum := 0
	for {
		res, ok := <-ch

		if !ok {
			fmt.Println("关闭通道")
			break
		}
		sum += res
	}
	return sum
}

func main() {

	ch := Producer()

	i := Consumer(ch)
	fmt.Println("----------")
	fmt.Println(i)
}
