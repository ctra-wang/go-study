package services

import (
	"bufio"
	"fmt"
	"os"
)

// 关于scanner的问题：https://blog.csdn.net/JineD/article/details/125435538

func ReadUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return "exit"
}

func ReadUserInputWithPrompt(prompt string) string {
	fmt.Print(prompt)
	return ReadUserInput()
}

func AskUserQuestion() string {
	return ReadUserInputWithPrompt("You: ")
}
