package cmd

import (
	"github.com/khytryy/taskig/internal/conf"
	"github.com/spf13/cobra"
)

var calendarTidy = &cobra.Command{
	Use:   "tidy",
	Short: "Delete all finished events",

	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return conf.TidyEvents()
	},
}

func init() {
	calendarCmd.AddCommand(calendarTidy)
}
