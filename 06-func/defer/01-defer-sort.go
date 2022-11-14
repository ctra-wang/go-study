package main

import "fmt"

// defer 经常处理一些资源的释放，文件关闭，解锁
func main() {
	fmt.Println("|-------------- Test01 func")

	// defer 先入后出
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")

	/*
		start...
		end...
		3
		2
		1
	*/
}
