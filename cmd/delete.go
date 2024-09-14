package cmd

import (
	"fmt"
	"os"
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
				fmt.Fprintf(os.Stderr, "Error: Invalid task ID: %v\n", err)
				return
			}

			err = task.DeleteTask(id, filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error deleting task: %v\n", err)
				return
			}

			fmt.Println("Task deleted successfully")
		},
	}
}
