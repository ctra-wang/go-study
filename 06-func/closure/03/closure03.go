package main

import "fmt"

func A() {
	fmt.Println("-----a")
}

func B() {
	// 1、defer 一定要在 panic前定義
	// 2、recover() 要配合defer 一起使用
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic(" panic b")
}

func C() {
	fmt.Println("-----c")
}

func main() {
	A()
	B()
	C()
}
