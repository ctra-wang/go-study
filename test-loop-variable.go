package main

import "fmt"

func main2() {
	//var prints []func()
	//for _, v := range []int{1, 2, 3} {
	//	temp := v
	//	prints = append(prints, func() { fmt.Println(temp) })
	//}
	//for _, print := range prints {
	//	print()
	//}

	//var prints []func()
	//for i := 1; i <= 3; i++ {
	//	prints = append(prints, func() { fmt.Println(i) })
	//}
	//for _, print := range prints {
	//	print()
	//}

	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}

}
