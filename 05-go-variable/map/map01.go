package _map

import "fmt"

func Test01() {
	fmt.Println("|-------------- Test01 map")
	// map类型 定义
	var map1 map[string]int
	fmt.Println(map1 == nil)
	// map类型 初始化
	map1 = make(map[string]int, 8)
	fmt.Println(map1 == nil)
	fmt.Printf("%T,len:%d \n", map1, len(map1))
}

func Test02() {
	fmt.Println("|-------------- Test02 map")

	// map类型 定义
	var map1 map[int]int
	// panic: assignment to entry in nil map
	// 如果不初始化则无法赋值
	//map1[2] = 100

	// map类型 初始化
	map1 = make(map[int]int)

	for i := 0; i < 19; i++ {
		map1[i] = i
	}
	fmt.Println(map1)
}

func Test03() {
	fmt.Println("|-------------- Test03 map")

	// map类型 初始化
	map1 := make(map[int]int)
	// 即使初始化长度不给，在遍历时都可以赋值
	for i := 0; i < 19; i++ {
		map1[i] = i
	}
	fmt.Println(map1)
}

// Test04 赋值操作，以及循环遍历
func Test04() {
	fmt.Println("|-------------- Test04 map")

	// map类型 初始化
	map1 := make(map[string]int)
	map1["num1"] = 100
	map1["num2"] = 200
	for v, k := range map1 {
		fmt.Printf("%s,%d \n", v, k)
	}
}

// Test05 delete()函数
func Test05() {
	fmt.Println("|-------------- Test05 map")

	// map类型 初始化
	map1 := make(map[string]int)
	map1["num1"] = 100
	map1["num2"] = 200
	map1["num3"] = 300

	for v, k := range map1 {
		fmt.Printf("%s,%d \n", v, k)
	}

	fmt.Println("------------删除了 key num2")
	delete(map1, "num2")
	for v, k := range map1 {
		fmt.Printf("%s,%d \n", v, k)
	}
}
