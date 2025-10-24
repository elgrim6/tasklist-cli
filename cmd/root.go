package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "MyCLI is a simple example command-line tool",
	Long:  `MyCLI is a demo CLI app built in Go using Cobra.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to MyCLI! Use --help to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
