package main

import (
	"fmt"
	"reflect"
)

type Per struct {
	name string `json:"name",ini:"name"`
	age  string `json:"age",ini:"age"`
}

func withReflect(x interface{}) {
	vf := reflect.ValueOf(x)
	tf := reflect.TypeOf(x)
	num := tf.NumField()
	for i := 0; i < num; i++ {
		fmt.Println(tf.Field(i))
	}

	numVf := vf.NumMethod()

	fmt.Println("---------------------")
	fmt.Println(numVf)

}

func main() {

	p1 := Per{
		name: "123",
		age:  "213",
	}

	withReflect(p1)

}
