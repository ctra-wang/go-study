package main

import "fmt"

/*
将err定义为string类型，然后把EOF设置为常量
*/
type err string

func (e err) Error2() string {
	fmt.Println("____+++++___")

	return string(e)
}

func (e err) Error() string {
	fmt.Println("_______")
	return string(e)
}

const EOF = err("EOF")

func main() {
	fmt.Println(EOF)
}
