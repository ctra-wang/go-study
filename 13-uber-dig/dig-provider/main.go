package main

import (
	"fmt"
	"go.uber.org/dig"
)

// DI IOC AOP

type A struct {
	b *B
}

type B struct {
	Name string
}

func NewA(b *B) *A {
	return &A{
		b: b,
	}
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
	_ = c.Provide(NewA) // bean
	_ = c.Provide(NewB) // bean

	//b := NewB()
	//a := NewA(b)
	// 注入 bean 获取器函数，并通过闭包的方式从中取出 bean
	var a *A
	c.Invoke(func(_a *A) {
		a = _a
	})
	fmt.Println(fmt.Sprintf("got a: %+v, got b: %+v", a, a.b))
}
