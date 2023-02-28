package main

import (
	"encoding/json"
	"fmt"
)

// 函数式选项模式（Functional Options Pattern）
type Book struct {
	Name  string
	Price string `json:"price,omitempty"`
}

type Option func(*Book)

func NewBook_(options ...Option) *Book {
	b := &Book{}
	for _, option := range options {
		option(b)
		marshal, _ := json.Marshal(b)
		fmt.Println(string(marshal))
	}
	return b
}

func WithPrice(price string) Option {
	return func(book *Book) {
		book.Price = price
	}
}

func WithName(name string) Option {
	return func(book *Book) {
		book.Name = name
	}
}

func main() {

	b := NewBook_(WithName("dsadas"), WithPrice("jjjj198877"))
	marshal, _ := json.Marshal(b)

	fmt.Println(string(marshal))
}
