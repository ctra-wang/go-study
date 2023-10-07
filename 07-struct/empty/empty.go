package main

import (
	"fmt"
	"time"
	"unsafe"
)

type EST struct {
}

func main() {
	// 一、基础
	// 空结构实例 和 空结构体变量 本质是一样的
	// 1、所有空结构体地址都是一样的
	// 2、大小都为0（最独特的）
	var a EST
	var b struct{}
	fmt.Printf("a address %p, size %d \n", &a, unsafe.Sizeof(a))
	fmt.Printf("b address %p, size %d \n", &b, unsafe.Sizeof(b))
	if a == b {
		fmt.Println("......")
	}

	// 二、应用场景一：充当 set
	set := make(map[string]interface{}, 10)
	set["A"] = EST{}
	set["B"] = struct{}{}

	fmt.Println(len(set))
	for k, _ := range set {
		fmt.Println(k)
	}

	// 三、应用场景二：协程阻塞
	// 1、一种是通过 waitGroup：通过调wait函数把本协程阻塞掉
	// 2、通过 time.Sleep()
	// 3、使用 管道 channel 读写阻塞

	// 下面展示为channel方式阻塞 main
	ctra := make(chan EST, 0)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("等待子协程执行完成！！！")
		ctra <- EST{}
	}()
	fmt.Println("---------")
	<-ctra

}
