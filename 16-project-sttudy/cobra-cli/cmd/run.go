/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/myapp/services"

	"github.com/spf13/cobra"
)

type Context struct {
	PreviousQuery string
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a gp3 chat Davinci bot",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取 gpt3秘钥
		keyMsg = services.InitKeyMag()
		key := keyMsg.GetKey(keyName) // 查看是否存在
		if key == "" {
			fmt.Println(`Don't find open-api key, please set your key first.
Find your key from https://beta.openai.com/account/api-keys.
Run command : go-chat key -s <your key>`)
			return
		}

		// 如果拿到gpt3，则初始化gpt3
		services.InitClient(key)
		if cmd.Flag("interactive").Value.String() == "true" {
			InteractiveMode()
		} else if cmd.Flag("prompt").Value.String() != "" {
			services.GetAnswer(cmd.Flag("prompt").Value.String())
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("interactive", "i", true, "interactive mode")
	runCmd.Flags().StringP("prompt", "p", "hello gopher", "prompt mode")
}

func InteractiveMode() {
	// 实例化一个 CacheHistory对象，加上过期时间和销毁时间
	history := services.NewCacheHistory()
	fmt.Print("Welcome to the GPT-3 chat bot. \nType 'exit' to quit, Type 'clear' to clear the context. \n")
	for {
		question := services.AskUserQuestion()
		if question == "exit" {
			break
		}
		if question == "clear" {
			history.ClearQACache()
			continue
		}
		cache, b := history.GetQACache()
		if b {
			question = cache + "/n" + question
		}

		reply, ok := services.GetAnswer(question)
		if ok {
			history.SetQACache(question, reply)
		}
	}
}
