package main

import (
	"fmt"
)

// 面试：打印s1是否更改了索引为1的值？
func main() {
	s1 := []int{1, 2, 3, 4}

	s2 := s1
	s2[1] = 99
	fmt.Println(s2)
	fmt.Println(s1) // 明显这是引用传递，更改了
}
