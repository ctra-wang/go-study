package array_slice

import "fmt"

/*
	通过append 删除切片
*/

func Test13() {
	slice1 := []string{"tom", "jack", "rose", "jim"}
	fmt.Printf("slice1---%p,%v,len:%d,cap:%d \n", slice1, slice1, len(slice1), cap(slice1))

	slice1 = append(slice1[1:2], slice1[3:]...)
	fmt.Printf("slice1---%p,%v,len:%d,cap:%d \n", slice1, slice1, len(slice1), cap(slice1))

}
