package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"net/http"
)

func morning(w http.ResponseWriter, req *http.Request) {
	fmt.Println("morning")
}

func evening(w http.ResponseWriter, req *http.Request) {
	fmt.Println("evening")
}

type Morning struct {
}

func (m Morning) Start() {
	http.HandleFunc("/morning", morning)
	http.ListenAndServe("localhost:6001", nil)
}

func (m Morning) Stop() {
	fmt.Println("Stop morning service...")
}

type Evening struct{}

func (e Evening) Start() {
	http.HandleFunc("/evening", evening)
	http.ListenAndServe("localhost:8081", nil)
}

func (e Evening) Stop() {
	fmt.Println("Stop evening service...")
}

func main() {
	// 方式一：直接启动
	fmt.Println("Start morning service...")
	var morning Morning
	morning.Start()
	defer morning.Stop()

	fmt.Println("Start evening service...")
	var evening Evening
	evening.Start()
	defer evening.Stop()

	// 方式二：改造第一种
	//var wg sync.WaitGroup
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("Start morning service...")
	//	var morning Morning
	//	defer morning.Stop()
	//	morning.Start()
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	fmt.Println("Start evening service...")
	//	var evening Evening
	//	defer evening.Stop()
	//	evening.Start()
	//}()
	//
	//wg.Wait()

	// 方式三：改造第二种
	group := service.NewServiceGroup()
	defer group.Stop()
	group.Add(Morning{})
	group.Add(Evening{})
	group.Start()

}
