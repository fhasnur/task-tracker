package cmd

import (
	"fmt"
	"os"
	"strings"
	"task-tracker/pkg/task"

	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list [status]",
		Short: "List all tasks. You can filter tasks by status",
		Run: func(cmd *cobra.Command, args []string) {
			status := "all"
			if len(args) > 0 {
				status = args[0]
			}

			tasks, err := task.ListTasks(status, filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error listing tasks: %v\n", err)
				return
			}

			fmt.Printf(
				"%-3s | %-30s | %-12s | %-16s | %-16s\n",
				"ID",
				"Description",
				"Status",
				"Created At",
				"Updated At",
			)
			fmt.Println(strings.Repeat("-", 90))

			for _, task := range tasks {
				fmt.Printf(
					"%-3d | %-30s | %-12s | %-16s | %-16s\n",
					task.Id,
					task.Description,
					task.Status,
					task.CreatedAt.Format("2006-01-02 15:04"),
					task.UpdatedAt.Format("2006-01-02 15:04"),
				)
			}
		},
	}
}
