package services

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
)

// 设定gpt参数
const (
	maxTokens   = 3000
	temperature = 0.7
	engine      = gpt3.TextDavinci003Engine
)

var client gpt3.Client

func InitClient(api string) {
	client = gpt3.NewClient(api)
}

var history = []string{}

func GetAnswer(question string) (reply string, ok bool) {
	fmt.Print("GPT3-达芬奇:")
	ok = false
	reply = ""
	i := 0
	err := client.CompletionStreamWithEngine(context.Background(), engine, gpt3.CompletionRequest{
		Prompt:      []string{question},
		MaxTokens:   gpt3.IntPtr(maxTokens),
		Temperature: gpt3.Float32Ptr(temperature),
	}, func(resp *gpt3.CompletionResponse) {
		if i > 1 {
			fmt.Print(resp.Choices[0].Text)
			reply += resp.Choices[0].Text
		}
		i++
	})
	if err != nil {
		return "", false
	}

	if len(reply) > 0 {
		ok = true
	}
	fmt.Println()
	return reply, ok
}

func FormatQuestion(question string) string {
	return "Answer:" + question
}
