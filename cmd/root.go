package cmd

import "github.com/spf13/cobra"

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-cli",
		Short: "CLI to track and manage your tasks",
	}

	cmd.AddCommand(
		addCmd(),
		updateCmd(),
		deleteCmd(),
		markInProgressCmd(),
		markDoneCmd(),
	)

	return cmd
}

func Execute() error {
	return rootCmd().Execute()

}
