package main

import "fmt"

// 闭包 = 函数+外层变量引用

func hello(name string) func() {
	return func() {
		fmt.Println("hello ", name)
	}
}

func acl(types string) func(name string) string {
	return func(name string) string {
		// 使用了外层得变量 types
		return fmt.Sprintf("%s%s", name, types)
	}
}

func main() {
	res := hello("wangdada")
	res()

	res1 := acl(".avi")
	func1 := res1("wangdada")
	fmt.Println(func1)
}
