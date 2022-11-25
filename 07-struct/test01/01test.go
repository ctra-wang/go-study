package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {

	p1 := &student{
		name: "123",
		age:  123,
	}

	fmt.Printf("%p \n", p1)
	fmt.Printf("%p \n", &p1.name)
	fmt.Printf("%p \n", &p1.age)

	fmt.Println("===========================")
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		fmt.Printf("原始：%p \n", &stu)
		stu := stu // 这里如果不再次取值，则一直取得地址都为指向该结构体的指针
		fmt.Printf("新的：%p \n", &stu)
		m[stu.name] = &stu
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

}
