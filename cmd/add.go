package cmd

import (
	"fmt"
	"os"
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
			id, err := task.AddTask(description, filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error adding task: %v\n", err)
				return
			}
			fmt.Printf("Task added successfully (ID: %d)\n", id)
		},
	}
}
