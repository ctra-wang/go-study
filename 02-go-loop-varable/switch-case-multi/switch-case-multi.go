package main

import "fmt"

// 这里主要演示case多条件下的书写方式
func main() {
	switch temp := 16; temp % 2 {
	// 除以二取模为0，整除，为偶数
	case 0:
		fmt.Println("这是偶数")
	// 除以二取模为1，不能整除，为奇数
	case 1:
		fmt.Println("这是奇数")
	default:
		fmt.Println("啥也不是")
	}
	// question1：求数组[1, 3, 5, 7, 8]之和
	total := 0
	arr := []int{1, 3, 5, 7, 8}
	for _, temp := range arr {
		total += temp
		continue
	}
	fmt.Println(total)

	// question2：求数组 [1, 3, 5, 7, 8] 中2个元素之和等于8的所有坐标 即：[0 3] [1 2]
	arrDouble := [][]int{}
	for index, temp := range arr {
		for i := index + 1; i < len(arr); i++ {
			if temp+arr[i] == 8 {
				arrDouble = append(arrDouble, []int{index, i})
			}
		}
	}
	fmt.Println(arrDouble)
}
