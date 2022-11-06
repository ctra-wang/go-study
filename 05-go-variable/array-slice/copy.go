package array_slice

import "fmt"

func Test01() {
	fmt.Println("|-------------- Test01 使用数组创建 slice")

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 4, 4)
	fmt.Printf("slice1---%p,%v \n", slice1, slice1)
	fmt.Printf("slice2---%p,%v \n", slice2, slice2)
	// copy 将源切片按照索引赋值到新的切片上
	copy(slice2, slice1)
	slice2[2] = 998
	slice3 := slice1
	slice3[0] = 123
	fmt.Printf("slice1---%p,%v \n", slice1, slice1)
	fmt.Printf("slice2---%p,%v \n", slice2, slice2)
	fmt.Printf("slice3---%p,%v \n", slice3, slice3)
}

func Test12() {
	fmt.Println("|-------------- Test12 使用数组创建 slice")
	slice1 := []string{"1", "2", "3"}

	slice2 := make([]string, 3, 3)
	copy(slice2, slice1)
	fmt.Printf("slice2---%p,%v,len:%d,cap:%d \n", slice2, slice2, len(slice2), cap(slice2))

}
