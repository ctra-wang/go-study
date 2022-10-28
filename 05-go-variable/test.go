package main

import (
	"fmt"
	const_iota "go-test/05-go-variable/const-iota"
)

func main() {

	fmt.Println("01 |--------------- IOTA测试代码读取为：")
	fmt.Println(const_iota.N1, const_iota.N2, const_iota.N3, const_iota.N4, const_iota.N5)

	fmt.Println("02 |--------------- 读取的存储容量值为：")
	fmt.Println(const_iota.KB, const_iota.MB, const_iota.GB, const_iota.TB, const_iota.PB)

	fmt.Println("03 |--------------- IOTA测试代码读取为：")
	fmt.Println(const_iota.I1, const_iota.I2, const_iota.I3, const_iota.I4, const_iota.I5, const_iota.I6)

}
