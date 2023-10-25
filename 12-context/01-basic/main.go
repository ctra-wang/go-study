package main

import (
	"context"
	"fmt"
)

// context最基础用法
// 俩个空父级：

func Step1(ctx context.Context) context.Context {
	// 通过接收父级ctx，生成child
	child := context.WithValue(ctx, "name", "json")
	return child
}

func Step2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", 998)
	return child
}

// 读取context
func ReadCtx(ctx context.Context) {
	fmt.Printf("name %s \n", ctx.Value("name"))
	fmt.Printf("age %d \n", ctx.Value("age"))
	fmt.Println("-------------")
}

func main() {
	// 创建一个父级context（空 ctx）
	ctx := context.TODO()
	ReadCtx(ctx)
	ctx1 := Step1(ctx)
	ReadCtx(ctx1)
	ctx2 := Step2(ctx1)
	ReadCtx(ctx2)
}
