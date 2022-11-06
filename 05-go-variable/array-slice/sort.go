package array_slice

import (
	"fmt"
	"sort"
)

func Test14() {
	fmt.Println("|-------------- Test14 使用数组创建 slice")

	arr := [...]int{3, 7, 9, 4, 6, 1}
	// sort 对int类型数组进行排序
	// 这里数组转切片 arr[:]
	sort.Ints(arr[:])

	fmt.Println(arr)

}
