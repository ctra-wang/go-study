package lib2

import "fmt"

// Mode 定义一个常量
// 首字母大写，对外可见(可在其它包中使用)
const Mode = 2

// num 定义一个全局整型变量
// 首字母小写，对外不可见(只能在当前包内使用)
var num = 200

// 初始化函数 类似；默认无参构造函数
func init() {
	fmt.Println("lib2---------------init")
	// 说明 定义的变量、常量会优先加载
	fmt.Println(num)
	fmt.Println(Mode)
}

func Lib2Test() {
	fmt.Println("Lib22222Test")

}
