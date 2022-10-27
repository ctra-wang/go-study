package main

import "fmt"

func main() {
	var prints []func()
	for i := 1; i <= 3; i++ {
		// 如果此处不重新将变量赋值，则会发现打印的都为相同的内容
		// i := i   // 取消注释，则为正确输出
		prints = append(prints, func() { fmt.Println(i) })
	}

	for _, print := range prints {
		print()
	}
}
