package main

import "fmt"

// struct 是值类型的，go语言中没有构造函数，通过自定义方法可以实现
type person struct {
	age     int
	name    string
	address string
}

// NewPerson 模拟go-构造函数
// 1、为什么这里要返回 *person 而不是 person
// 因为struct是值类型的，如果结构体定义的过大，则每次（深）copy都会整体复制一片内存，使用传地址为最佳实践
func NewPerson(age int, name, address string) *person {
	return &person{
		age:     age,
		name:    name,
		address: address,
	}
}

func main() {
	// 实例化一个对象
	p := NewPerson(10, "wang", "tj")
	fmt.Printf("%#v", p)
}
