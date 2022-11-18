package main

import "fmt"

// struct嵌套 匿名
type Animals struct {
	name string
	age  int
}

func (a Animals) move() {
	fmt.Printf("%d的%s,正在飞奔中。。。\n", a.age, a.name)
}

func (a Animals) jump() {
	fmt.Printf("%d的%s,正在跳跃中。。。\n", a.age, a.name)
}

type dog struct {
	legs int
	*Animals
}

func NewDog(legs int, animal *Animals) *dog {
	return &dog{
		legs:    legs,
		Animals: animal,
	}
}

func main() {

	test := NewDog(4, &Animals{age: 12, name: "wangwandui"})
	test.move()
	test.jump()
}
