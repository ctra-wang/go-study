package main

import "fmt"

// 结论：结构体占用一块连续的内存。

type arrs struct {
	a1 int8
	a2 int8
	a3 int8
	a4 int8
}

func main() {
	test := arrs{
		1, 2, 3, 4,
	}

	fmt.Printf("a1 %p \n", &test.a1)
	fmt.Printf("a2 %p \n", &test.a2)
	fmt.Printf("a3 %p \n", &test.a3)
	fmt.Printf("a4 %p \n", &test.a4)

	/* 输出结果
	a1 0xc0000a6058
	a2 0xc0000a6059
	a3 0xc0000a605a
	a4 0xc0000a605b
	*/
}
