package main

import (
	"fmt"
)

func main() {

	s1 := []int{1, 2, 3, 4}

	s2 := make([]int, len(s1))
	res := copy(s2, s1) // 使用copy返回拷贝的元素数量
	fmt.Println(res)
	s2[1] = 99
	fmt.Println(s2)
	fmt.Println(s1) // 明显这是引用传递，更改了
}
