package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name",ini:"s_name"`
	age  int    `json:"age",ini:"s_age"`
}

type Test interface {
	//dosomething()
}

func (p Person) Profile() string {
	msg := fmt.Sprintf("你好我是：%s，我今年：%d 岁啦", p.Name, p.age)
	fmt.Println(msg)
	return msg
}

func (p Person) Sleep(arg string) string {
	msg := fmt.Sprintf("%s 需要在%s前赶紧睡觉", p.Name, arg)
	fmt.Println(msg)
	return msg
}

// 根据反射 获取结构体方法
func printMethod(x interface{}) {
	tf := reflect.TypeOf(x)
	vf := reflect.ValueOf(x)

	// 如果方法名小写，则无法取到该方法
	fmt.Println(tf.NumMethod())
	fmt.Println(vf.NumMethod())

	for i := 0; i < vf.NumMethod(); i++ {
		method := vf.Method(i)
		fmt.Printf("Name:%v ,type:%v  \n", tf.Method(i).Name, method.Type())

		args := []reflect.Value{}
		if tf.Method(i).Name == "Sleep" {
			args = append(args, reflect.ValueOf("20:00点"))
		}
		method.Call(args)
	}

}

// 通过反射获取结构体：
// 重要函数：ref := reflect.TypeOf(p1)、ref.NumField()、ref.Field(i)
func main() {

	// ------------------------ 通过反射获取结构内容 ------------------------
	p1 := Person{
		Name: "wangdada",
		age:  10,
	}

	ref := reflect.TypeOf(p1)
	fmt.Printf("%v %v \n", ref.Name(), ref.Kind())

	// 使用 ref.NumIn() + ref.Field() 遍历所有字段
	// ref.NumIn() ：返回元素的个数
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		fmt.Printf("Name:%v ,type:%v ,Tag:%v \n", field.Name, field.Type, field.Tag)
	}

	// 方式二
	obj, ok := ref.FieldByName("Age")
	if ok {
		fmt.Printf("Name:%v ,type:%v ,Tag:%v \n", obj.Name, obj.Type, obj.Tag)
	}

	// ------------------------ 通过反射获取结构方法 ------------------------
	fmt.Println("------------------------------")
	var test Test
	test = p1
	printMethod(test)

}
