package main

import "fmt"

func main() {
	var prints []func()
	for _, v := range []int{1, 2, 3} {
		// 如果此处不重新将变量赋值，则会发现打印的都为相同的内容
		// v := v   // 取消注释，则为正确输出
		prints = append(prints, func() { fmt.Println(v) })
	}
	for _, print := range prints {
		print()
	}
}
