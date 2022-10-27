package lib1

import "fmt"

// 包级别标识符的可见性

// Mode 定义一个常量
// 首字母大写，对外可见(可在其它包中使用)
const Mode = 1

// num 定义一个全局整型变量
// 首字母小写，对外不可见(只能在当前包内使用)
var num = 100

// 初始化函数 类似；默认无参构造函数
func init() {
	fmt.Println("lib1---------------init")
	// 说明 定义的变量、常量会优先加载
	fmt.Println(num)
	fmt.Println(Mode)
}

// Lib1Test 只是为了测试一下方法提示语
func Lib1Test() {
	fmt.Println("Lib11111Test")

}
