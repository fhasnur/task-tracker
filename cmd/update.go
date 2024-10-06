package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fhasnur/task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func updateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update [id] [new description]",
		Short: "Update an existing task",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Invalid task ID: %v\n", err)
				return
			}

			newDescription := args[1]
			err = task.UpdateTask(id, newDescription, filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error updating task: %v\n", err)
				return
			}

			fmt.Println("Task updated successfully")
		},
	}
}
