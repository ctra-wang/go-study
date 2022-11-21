package main

import "fmt"

// 空 interface类型使用场景：
// 一般使用不确定的参数
// 需要配合断言使用
func main() {
	var x interface{}

	x = "wangdada"
	fmt.Println(x)
	x = 190
	fmt.Println(x)
	x = false
	fmt.Println(x)

	b, ok := x.(bool)
	if !ok {
		fmt.Println("不是类型:bool:", b)
	} else {
		fmt.Println("是类型:bool:", b)
	}

	switch v := x.(type) {
	case string:
		fmt.Println("string:", v)
	case bool:
		fmt.Println("bool:", v)
	case int:
		fmt.Println("int:", v)
	}
}
