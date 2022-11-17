package main

import "fmt"

type Person struct {
	age     int8
	name    string
	address string
}

func NewPerson(age int8, name, address string) *Person {
	return &Person{
		age:     age,
		name:    name,
		address: address,
	}
}

// ----------------- 方法 -----------------
func (p *Person) setName(name string) {
	p.name = name
}

// 由于 struct 为值拷贝，无法修改
func (p Person) setName1(name string) {
	p.name = name
}

func (p *Person) setAge(age int8) {
	p.age = age
}

func main() {
	p1 := NewPerson(18, "wangdada", "jjj")

	fmt.Println(p1.name)
	p1.setName1("zhangsanfeng")
	fmt.Println(p1.name)
	// 正确的方式
	p1.setName("zhangsanfeng")
	p1.setAge(71)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

}
