package main

import (
	"fmt"
	go_array "go-test/05-go-variable/array-slice"
	byte_rune "go-test/05-go-variable/byte-rune"
	const_iota "go-test/05-go-variable/const-iota"
)

func init() {
	fmt.Println("|-------------- hhhhhhhhhhhhhhhhhhhh-----------------")

}

func main() {

	fmt.Println("01 |--------------- IOTA测试代码读取为：")
	fmt.Println(const_iota.N1, const_iota.N2, const_iota.N3, const_iota.N4, const_iota.N5)

	fmt.Println("02 |--------------- 读取的存储容量值为：")
	fmt.Println(const_iota.KB, const_iota.MB, const_iota.GB, const_iota.TB, const_iota.PB)

	fmt.Println("03 |--------------- IOTA测试代码读取为：")
	fmt.Println(const_iota.I1, const_iota.I2, const_iota.I3, const_iota.I4, const_iota.I5, const_iota.I6)

	fmt.Println(" |--------------- 分割线 ---------------|")
	byte_rune.Test()

	fmt.Println(" |--------------- 分割线 ---------------|")
	go_array.Test02()
	go_array.Test03()
	go_array.Test04()
	go_array.Test05()
	go_array.Test06()
	go_array.Test10()
	arr := []int{-3, -2, -1, 0, 1, 2, 7, 9, 11}
	res := go_array.Test11(arr)
	fmt.Println(res)
}
