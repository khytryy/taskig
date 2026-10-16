package cmd

import (
	"fmt"
	"time"

	"github.com/khytryy/taskig/internal/conf"
	"github.com/spf13/cobra"
)

var calendarSchedule = &cobra.Command{
	Use:     "schedule <name> <begin> <end>",
	Short:   "Schedule an event",
	Example: "schedule \"Math class\" \"2026-Oct-16 15:10\" \"2026-Oct-16 16:00\"",

	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		begin, err := time.ParseInLocation("2006-Jan-02 15:04", args[1], time.Local)
		if err != nil {
			return err
		}

		end, err := time.ParseInLocation("2006-Jan-02 15:04", args[2], time.Local)
		if err != nil {
			return err
		}

		conf.AddEvent(conf.Event{
			Begin: begin,
			End:   end,
			Name:  args[0],
		})

		fmt.Printf("added event %q\n", args[0])
		return nil
	},
}

func init() {
	calendarCmd.AddCommand(calendarSchedule)
}
