package main

import "fmt"

// person
type person struct {
	name string
}

func (p person) say() {
	fmt.Printf("%s 在大叫 \n", p.name)
}

// animal
type animal struct {
	name string
}

func (a animal) say() {
	fmt.Printf("%s 在大叫 \n", a.name)
}

// interface
type action interface {
	say()
}

//	 场景： 我们有俩个方法都实现了 say()
//		这里可以理解为设计模式：模板方法
//		这里需要一个方法，可以让他们任何一个结构体作为参数进行传入
//		则使用 action这个结构体作为参数
func templateFunc(a action) {
	a.say()
}

func main() {
	p1 := person{
		name: "wangdada",
	}

	a1 := animal{
		name: "xiaoniao",
	}
	templateFunc(p1)
	templateFunc(a1)
}
