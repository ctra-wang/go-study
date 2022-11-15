package main

import "fmt"

func main() {
	go func() {
		fmt.Println("------1")
	}()

	go func() {
		fmt.Println("------2")
	}()

	go func() {
		fmt.Println("------3")
	}()

	go func() {
		panic("asd")
		//fmt.Println("------1")
	}()

	go func() {
		fmt.Println("------4")
	}()

	go func() {
		fmt.Println("------6")
	}()
}
