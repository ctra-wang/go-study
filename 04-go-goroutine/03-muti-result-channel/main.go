package main

import "fmt"

// 当向通道中发送完数据时，我们可以通过close函数来关闭通道。
// 当一个通道被关闭后，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值。

// 对一个通道执行接收操作时支持使用如下多返回值模式
func recv(c chan int) {
	for {
		// 通道内的值被接收完后再对通道执行接收操作得到的值会一直都是对应元素类型的零值。那我们如何判断一个通道是否被关闭了呢？
		// A：ok：通道ch关闭时返回 false，否则返回 true。
		res, ok := <-c
		if !ok {
			fmt.Println("通道关闭")
			break
		}
		fmt.Println(res)
	}
	//fmt.Println("接收成功：", res)
}
func main() {

	chan1 := make(chan int, 2)

	chan1 <- 10
	chan1 <- 20

	close(chan1)
	recv(chan1)
	//fmt.Println("发送成功")
}
