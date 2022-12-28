package _4_go_goroutine

import (
	"fmt"
	"time"
)

func goFunc(i int) {
	fmt.Println("goroutine ", i, "...")
}

func main1() {
	for i := 0; i < 100; i++ {
		go goFunc(i)
	}

	time.Sleep(time.Second)
}
