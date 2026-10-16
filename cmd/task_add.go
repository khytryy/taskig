package cmd

import (
	"fmt"
	"taskig/internal/conf"

	"github.com/spf13/cobra"
)

var taskAdd = &cobra.Command{
	Use:   "add <name>",
	Short: "Create a new todo task",

	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := conf.AddTodoTask(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("added %q\n", args[0])
		return nil
	},
}

func init() {
	taskCmd.AddCommand(taskAdd)
}
