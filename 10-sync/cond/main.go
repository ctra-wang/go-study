package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	race()
}

// sync.Cond 三个方法：
// 1、Wait 阻塞当前携程
// 2、Signal,唤醒一个等待时间最长的协程
// 3、BroadCast ，唤醒所有等待的协程

func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			cond.Wait() //等待发令枪响
			fmt.Println(num, "号开始跑....")
			cond.L.Unlock()
		}(i)
	}

	// 等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() // 发令枪响
	}()
	wg.Wait()
}
