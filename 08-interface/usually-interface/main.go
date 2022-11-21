package main

import (
	"fmt"
)

type person struct {
	name string
}

type action interface {
	say()
}

func (p person) say() {
	fmt.Printf("%s 在大叫 \n", p.name)
}

func main() {
	var a action
	var a2 action
	p1 := person{
		name: "wangdada",
	}

	p2 := &person{
		name: "zhangsanfeng",
	}
	a = p1
	a2 = p2
	a.say()
	a2.say()
}
