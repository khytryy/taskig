package cmd

import (
	"fmt"
	"taskig/internal/conf"

	"github.com/spf13/cobra"
)

var taskComplete = &cobra.Command{
	Use:   "complete <name>",
	Short: "Delete todo task",

	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := conf.RmTodoTask(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("completed %q\n", args[0])
		return nil
	},
}

func init() {
	taskCmd.AddCommand(taskComplete)
}
