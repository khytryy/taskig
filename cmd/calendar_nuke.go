package cmd

import (
	"fmt"

	"github.com/khytryy/taskig/internal/conf"
	"github.com/spf13/cobra"
)

var calendarNuke = &cobra.Command{
	Use:   "nuke",
	Short: "Deletes all tasks",

	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := conf.RmEvents()
		if err != nil {
			return err
		}

		fmt.Println("nuked all events")
		return nil
	},
}

func init() {
	calendarCmd.AddCommand(calendarNuke)
}
