package cmd

import "github.com/spf13/cobra"

var taskCmd = &cobra.Command{
	Use:   "task [OPTIONS]",
	Short: "Manage a todo task",
}

func init() {
	rootCmd.AddCommand(taskCmd)
}
