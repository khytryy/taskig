package cmd

import "github.com/spf13/cobra"

var calendarCmd = &cobra.Command{
	Use:   "calendar [OPTIONS]",
	Short: "Manage the calendar",
}

func init() {
	rootCmd.AddCommand(calendarCmd)
}
