package _map

import (
	"fmt"
	"math/rand"
	"sort"
)

// Test06  证明切片是无序的：下面可以看到将一个切片赋值，然后循环打印
func Test06() {
	fmt.Println("|-------------- Test06 map")
	slice := make(map[string]int, 100)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%d", i)
		slice[key] = rand.Intn(99) // 设置0~99的一个随机数
	}

	for key, temp := range slice {
		fmt.Printf("%s,%d \n", key, temp)
	}
}

// Test07  证明切片是无序的：下面可以看到将一个切片赋值，然后循环打印
func Test07() {
	fmt.Println("|-------------- Test07 map")
	slice := make(map[string]int, 100)

	for i := 0; i < 50; i++ {
		// todo 这里无底深坑  stu%d 这样写法 导致sory排序失败
		//key := fmt.Sprintf("stu%d", i)
		// 正确写法：就是1前面带着0 为01
		key := fmt.Sprintf("stu%02d", i)
		slice[key] = rand.Intn(99) // 设置0~99的一个随机数
	}

	//for key, temp := range slice {
	//	fmt.Printf("%s,%d \n", key, temp)
	//}

	// 将切片根据key的顺序进行排序
	// 1、将keys新创建一个新的切片
	// 2、将keys进行排序
	// 3、根据keys循环打印
	keys := make([]string, 0)
	for key, _ := range slice {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	fmt.Println("--------------------------")
	sort.Strings(keys)

	sliceNew := make(map[string]int, 100)
	for _, temp := range keys {

		sliceNew[temp] = slice[temp]
		fmt.Printf("%s,%d \n", temp, slice[temp])
	}

}
