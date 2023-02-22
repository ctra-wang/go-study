package main

import "fmt"

type Circle struct {
	Radius float64
}

type Sharper interface {
	Area() float64
	Scale(factor float64)
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Scale(factor float64) {
	c.Radius = c.Radius * factor
}

func main() {

	// common sense
	// 1、value methods can be invoked on pointers and values
	// 2、pointer methods can only be invoked on pointers
	// 3、when the value is addresssable,you can winvoking apointer method on a value by insrting the address operator automatically

	// 1、
	c := &Circle{20}
	c.Scale(20)
	fmt.Println(c.Area())

	// 3
	circle := Circle{18}
	circle.Scale(2) // equals to：(&circle).Scale(2)
	circle.Area()   // equals to：(&circle).Scale(2)
}
