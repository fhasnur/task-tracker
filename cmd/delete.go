package cmd

import (
	"fmt"
	"strconv"
	"task-tracker/pkg/task"

	"github.com/spf13/cobra"
)

func deleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID:", err)
				return
			}

			err = task.DeleteTask(id, filename)
			if err != nil {
				fmt.Println("Error deleting task:", err)
				return
			}
			fmt.Println("Task deleted successfully")
		},
	}
}
