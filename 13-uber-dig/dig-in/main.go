package main

import (
	"fmt"
	"go.uber.org/dig"
)

// DI IOC AOP

type A struct {
	dig.In
	B *B
}

type B struct {
	Name string
}

func NewB() *B {
	return &B{
		Name: "i am b",
	}
}

func main() {
	// 创建一个容器 container  dig
	c := dig.New()

	// 注入各个bean的构造函数
	_ = c.Provide(NewB) // bean

	//b := NewB()
	//a := NewA(b)
	// 使用 bean A 的 struct 形式，与 container 进行 Invoke 交互
	var a A
	c.Invoke(func(_a A) {
		a = _a
	})

	fmt.Println(fmt.Sprintf("got a: %+v, got b: %+v", a, a.B))
}
