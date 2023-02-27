/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/myapp/services"
	"os"

	"github.com/spf13/cobra"
)

const (
	keyName = "chat-go-key"
	Version = "0.0.1"
	AppName = "chat-go"
)

var keyMsg *services.KeyMag

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   AppName,
	Short: "CLI ChatBot Power By Gpt3",
	Run: func(cmd *cobra.Command, args []string) {
		if flag := cmd.Flag("version"); flag != nil && flag.Value.String() == "true" {
			cmd.Println(AppName, "v"+Version)
		} else {
			cmd.Help()
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	services.NewKeyMag()
	rootCmd.Flags().BoolP("version", "v", false, "version of "+AppName)
}
