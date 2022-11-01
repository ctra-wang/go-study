package array_slice

import "fmt"

func change(arr [3]int) {
	arr[0] = 998
}

// Test02 判断 array数组 为值类型
func Test02() {
	fmt.Println("|-------------- Test02 使用数组创建 slice")
	arr := [3]int{1, 2, 3}
	fmt.Printf("%T \n", arr)
	fmt.Println(arr)
	change(arr)
	fmt.Println(arr)
}

// Test03 使用make创建slice
func Test03() {
	fmt.Println("|-------------- Test03 使用数组创建 slice")
	// 初始长度位5，总容量为10
	arr2 := make([]int, 5, 10)
	fmt.Println(arr2)
	fmt.Println(len(arr2))
	fmt.Println(cap(arr2)) // 如果不设置容量，则默认与长度相同
	fmt.Printf("%T \n", arr2)
}

// Test04 使用数组创建 slice
func Test04() {
	fmt.Println("|-------------- Test04 使用数组创建 slice")
	arr2 := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1 := arr2[:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1)) // 如果不设置容量，则默认与长度相同

	slice2 := arr2[3:6]
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2)) // 如果不设置容量，则默认与长度相同
}

// Test05 判断slice是否为nil
func Test05() {
	fmt.Println("|-------------- Test05 使用数组创建 slice")
	var slice1 []int     // 声明 不初始化
	var slice2 = []int{} // 声明 初始化
	var slice3 = make([]int, 0)

	if slice1 == nil {
		fmt.Println("!============= slice1 == nil")
	}
	fmt.Println(slice1, cap(slice1), len(slice1))
	if slice2 == nil {
		fmt.Println("!============= slice2 == nil")
	}
	fmt.Println(slice2, cap(slice2), len(slice2))
	if slice3 == nil {
		fmt.Println("!============= slice3 == nil")
	}
	fmt.Println(slice3, cap(slice3), len(slice3))
}

// Test06 判断slice 为引用类型
func Test06() {
	fmt.Println("|-------------- Test06 使用数组创建 slice")
	// 通过引用类型，可以看到 通过修改 slice2同时也将slice1得值进行了修改，说明他们指向了相同得指针
	var slice1 = make([]int, 3)
	slice2 := slice1
	slice2[0] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}
