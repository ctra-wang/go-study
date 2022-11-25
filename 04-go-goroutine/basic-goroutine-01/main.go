package main

import (
	"fmt"
	"sync"
)

// 声明全局等待组变量
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello")
	wg.Done() // 协程 执行完成
}

func main() {

	wg.Add(2) // 这里要和执行的协程数保持一致，如果不一致则 deadlock，如果为 2，则需要再加一个协程
	go hello()

	wg.Wait() // 阻塞，等待全部协成完成

	fmt.Println("main hello")
}
