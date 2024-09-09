package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("Task Tracker")
	var rootCmd = &cobra.Command{
		Use:   "task-cli [task]",
		Short: "CLI to track and manage your tasks",
		Run: func(cmd *cobra.Command, args []string) {
			task := args[0]
			fmt.Println(task)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
