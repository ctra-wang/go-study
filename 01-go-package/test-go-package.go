package main

import (
	"fmt"
	"go-test/01-go-package/lib1"
	"go-test/01-go-package/lib2"
	_ "go-test/01-go-package/lib3" // 引入lib3，但是不使用，则前面加入 _ 作为匿名，然后会读取他的init()
)

func main() {
	lib1.Lib1Test()
	fmt.Println("----------------- 华丽的分割线 -----------------")
	lib2.Lib2Test()

	// import . "go-test/01-go-package/lib3" 可以直接使用包内方法
	// Lib3Test()
}
