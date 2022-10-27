package main

import (
	"fmt"
	"go-test/go-init/lib1"
	"go-test/go-init/lib2"
)

func main() {
	lib1.Lib1Test()
	fmt.Println("----------------- 华丽的分割线 -----------------")
	lib2.Lib2Test()
}
