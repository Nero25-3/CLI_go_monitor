package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check if the monitor process is running",
	RunE: func(cmd *cobra.Command, args []string) error {
		pidFile := "monitor.pid"
		data, err := os.ReadFile(pidFile)
		if err != nil {
			fmt.Println("Monitor is NOT running.")
			return nil
		}
		fmt.Println("Monitor is running with PID", string(data))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
