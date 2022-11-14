package main

import "fmt"

var num = 10

func hello() {
	num := 100
	fmt.Println("hello ", num)
}

func main() {
	hello()
}
