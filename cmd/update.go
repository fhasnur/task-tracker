package cmd

import (
	"fmt"
	"strconv"
	"task-tracker/pkg/task"

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
				fmt.Println("Invalid task ID:", err)
				return
			}

			newDescription := args[1]
			err = task.UpdateTask(id, newDescription, filename)
			if err != nil {
				fmt.Println("Error updating task:", err)
				return
			}
			fmt.Println("Task updated successfully")
		},
	}
}
