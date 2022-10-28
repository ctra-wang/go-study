package byte_rune

import "fmt"

func Test() {

	// byte uint8 别名 ASCII码
	// rune int32 别名

	var c1 byte = 'c'
	var c2 rune = 'c'

	fmt.Println(c1, c2)
	// 通过 %T 来判断参数类型
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	str := "wangdada王大大"
	length := GetStringSize(str)
	fmt.Println(length)

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println("")
	for _, r := range str {
		fmt.Printf("%c ", r)
	}

}
