package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type EIF interface {
}

// 1、任何类型都属于空接口
func test(arg EIF) {

}

func main() {

	var e1 interface{}
	var e2 EIF
	var e3 any // 1.18版本以后 官方建议使用

	fmt.Printf(".......address:%p,size:%d\n", &e1, unsafe.Sizeof(e1))
	fmt.Printf(".......address:%p,size:%d\n", &e2, unsafe.Sizeof(e2))
	fmt.Printf(".......address:%p,size:%d\n", &e3, unsafe.Sizeof(e3))

	test(10)
	test("10")

	e2 = 10
	e2 = "10"

	fmt.Println(e2)
	fmt.Println(reflect.TypeOf(e2).Name())
}
