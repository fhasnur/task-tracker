package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fhasnur/task-tracker/internal/task"

	"github.com/spf13/cobra"
)

func markInProgressCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mark-in-progress [id]",
		Short: "Mark a task as in-progress",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Invalid task ID: %v\n", err)
				return
			}

			err = task.MarkTask(id, "in-progress", filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error marking task as in-progress: %v\n", err)
				return
			}

			fmt.Println("Task marked as in-progress successfully")
		},
	}
}

func markDoneCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mark-done [id]",
		Short: "Mark a task as done",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Invalid task ID: %v\n", err)
				return
			}

			err = task.MarkTask(id, "done", filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error marking task as done: %v\n", err)
				return
			}

			fmt.Println("Task marked as done successfully")
		},
	}
}
