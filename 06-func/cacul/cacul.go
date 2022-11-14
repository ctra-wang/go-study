package main

import "fmt"

func add(a int, b int) int {
	return (a + b)
}

func sub(a int, b int) int {
	return (a - b)
}

// 第三个参数 cal func(int, int) int：由于add方法正好满足这个形式，可以将add作为参数传入
func cacul(x int, y int, cal func(int, int) int) int {
	return cal(x, y)
}

func main() {
	res := cacul(1, 2, add)
	fmt.Println(res)

	res1 := cacul(1, 2, sub)
	fmt.Println(res1)
}
