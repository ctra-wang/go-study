package main

import (
	"bytes"
	"fmt"
	"io"
)

/*
IO 包的设计缺陷：EOF 使用var定义，而不是const，允许被改变
*/

func init() {
	// 读取文件结尾，如果为nil 则一直会无限读取文件，导致下面会timeout
	io.EOF = nil
}

func main() {

	reader := bytes.NewReader([]byte{1, 2, 3})
	res, err := io.ReadAll(reader)
	fmt.Println(res, err) // 死循环 timeout
}
