package main

import "fmt"

// 匿名函数
func main() {
	// 方式一：定义、执行
	// 定义一个匿名函数 给一个变量
	test01 := func() {
		fmt.Println("hello")
	}
	test01() //执行此函数

	// 方式二：定义立即执行
	func() {
		fmt.Println("hello")
	}()

}
