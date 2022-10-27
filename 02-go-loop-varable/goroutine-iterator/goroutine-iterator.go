package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		// 如果此处不重新将变量赋值，则会发现打印的都为相同的内容
		// v := v  // 取消注释，则为正确输出
		go func() {
			// 在这里
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}
