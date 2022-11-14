package _map

import (
	"fmt"
	"strings"
)

func Test08() {
	fmt.Println("|-------------- Test08 map")

	// 1、声明 此时得 mapSlice 所有得元素都为nil [nil nil nil nil nil nil nil]
	var mapSlice = make([]map[string]int, 8, 8)
	// 证明 当前 [nil nil nil nil nil nil nil]
	fmt.Println(mapSlice[0] == nil)

	// 2、初始化 才能使用
	mapSlice[0] = make(map[string]int)
	fmt.Println(mapSlice[1] == nil)
	fmt.Println(mapSlice[0] == nil)

	//for i := 0; i < 8; i++ {
	//	mapSlice[i]["gff"] = 100
	//}

	mapSlice[0]["gff"] = 100
	//mapSlice[1]["123"] = 100
	//mapSlice[2]["123"] = 100
	fmt.Println(mapSlice)
}

func Test09() {
	fmt.Println("|-------------- Test09 map")
	var sliceMap = make(map[string][]int, 8)

	v, ok := sliceMap["good"]
	if ok {
		fmt.Println(v)
	} else {
		sliceMap["good"] = make([]int, 8)
		sliceMap["good"][0] = 100
		sliceMap["good"][1] = 200
		sliceMap["good"][3] = 100
	}

	fmt.Println(sliceMap)
}

func Test10() {
	const name = "how do you do"
	arr := strings.Split(name, " ")

	var res = make(map[string]int, 8)
	for _, temp := range arr {
		v, ok := res[temp]
		fmt.Println(v)
		if ok {
			res[temp] += 1
		} else {
			res[temp] = 1
		}
	}
	for k, v := range res {
		fmt.Println(k, v)
	}
}
