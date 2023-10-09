package main

import "fmt"

func test01() (result int) {
	defer func() {
		result++
	}()
	return
}

func test02() (result int) {
	result = 0 //return语句不是一条原子调用，return xxx其实是赋值＋ret指令
	func() {   //defer被插入到return之前执行，也就是赋返回值和ret指令之间
		result++
	}()
	return
}

func main() {
	fmt.Println(test01())
}
