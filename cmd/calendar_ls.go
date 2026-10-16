package cmd

import (
	"github.com/khytryy/taskig/internal/conf"
	"github.com/spf13/cobra"
)

var all bool

var calendarLs = &cobra.Command{
	Use:   "ls",
	Short: "List all events",
	Long:  "List all events in time order",

	Args: cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if all {
			return conf.ListAllEvents()
		}
		return conf.ListEvents()
	},
}

func init() {
	calendarLs.Flags().BoolVarP(&all, "all", "a", false, "list events including completed ones")
	calendarCmd.AddCommand(calendarLs)
}
