package cmd

import (
	"fmt"
	"strings"
	"task-tracker/pkg/task"

	"github.com/spf13/cobra"
)

func listAllCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := task.ListAllTasks(filename)
			if err != nil {
				fmt.Println("Error listing tasks:", err)
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
