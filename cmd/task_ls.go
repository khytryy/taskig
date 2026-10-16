package cmd

import (
	"taskig/internal/conf"

	"github.com/spf13/cobra"
)

var taskLs = &cobra.Command{
	Use:   "ls",
	Short: "List all todos",

	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := conf.ListTodos()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	taskCmd.AddCommand(taskLs)
}
