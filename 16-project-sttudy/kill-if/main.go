package main

import "fmt"

// kill-if 配合说明文章地址：https://ctraplatform.blog.csdn.net/article/details/129105030

// SomeConditions 我是一个单纯的结构体
type SomeConditions struct {
	Price      float64
	OrderCount int
	TotalCount int
	MemberShip int
}

// 传统的if-else判断导致业务逻辑一直在添加if-else上不断耦合
func main2() {
	var a = SomeConditions{
		Price:      1.0,
		OrderCount: 1,
		TotalCount: 20,
		MemberShip: 1,
	}
	if a.Price > 0 {
		println("true")
	}
	if a.TotalCount > a.OrderCount {
		println("true")
	}
	if a.MemberShip == 1 {
		println("true")
	}
	if a.Price < 100 && a.MemberShip == 2 {
		println("true")
	}
}

// Rule 规则接口，抛出统一判断方法，各自实现
type Rule interface {
	Check(s *SomeConditions) bool
}

// 聚合所有条件（责任链）类似一个拦截器
// 这里记住一定要返回错误的，即：满足都是通过的，只拦截不满足的
func chain(s *SomeConditions, rules ...Rule) bool {
	for _, rule := range rules {
		if !rule.Check(s) {
			return false
		}
	}
	return true
}

func main() {
	SomeConditions := &SomeConditions{
		Price:      1.00,
		OrderCount: 2,
		TotalCount: 10,
		MemberShip: 2,
	}
	rules := []Rule{
		&PriceCheck{},
		&OrderCountCheck{},
		&MemberShipCheck{},
		&DiscountCheck{},
	}
	// 进行判断
	flag := chain(SomeConditions, rules...)
	fmt.Println(flag)
}

// ------------		下面是4种业务的判断	------------

type PriceCheck struct{}

func (p *PriceCheck) Check(s *SomeConditions) bool {
	return s.Price > 0
}

type OrderCountCheck struct{}

func (o *OrderCountCheck) Check(s *SomeConditions) bool {
	return s.TotalCount > s.OrderCount
}

type MemberShipCheck struct{}

func (m *MemberShipCheck) Check(s *SomeConditions) bool {
	return s.MemberShip == 2
}

type DiscountCheck struct{}

func (d *DiscountCheck) Check(s *SomeConditions) bool {

	return s.Price > 10 && s.MemberShip == 2
}
