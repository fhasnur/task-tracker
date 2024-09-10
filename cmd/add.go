package cmd

import (
	"fmt"
	"task-tracker/pkg/task"

	"github.com/spf13/cobra"
)

const filename = "tasks.json"

func addCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [description]",
		Short: "Add a new task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			description := args[0]
			err := task.AddTask(description, filename)
			if err != nil {
				fmt.Println("Error adding task:", err)
				return
			}
			fmt.Println("Task added succesfully")
		},
	}
}
